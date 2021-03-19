package main

import (
	"flag"
	"fmt"
	"hash/crc32"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"syscall"
	"time"
)

const count = 1000

type paths []string

type fileStat struct {
	hash []byte
	size int64
	path string
}

func main() {
	var (
		//wg = &sync.WaitGroup{}
		ch = make(chan fileStat, count)
		path string
		clear bool
	)

	log.SetFlags(log.LUTC | log.Lmicroseconds | log.Lshortfile | log.Ldate)

	flag.Usage = func() {
		_, err := fmt.Fprintf(os.Stderr, "Usage of %s: This program searches for duplicate files in a folder and subfolders " +
			"and, if necessary, deletes them\n", filepath.Base(os.Args[0]))
		if err != nil {
			log.Println("error:", err)
		}
		flag.PrintDefaults()
	}
	flag.StringVar(&path,"folder", "", "required parameter. path to a folder.")
	flag.BoolVar(&clear,"clear", false,"bool parameter, default = false. if set to true - duplicates will be removed.")
	flag.Parse()

	if path == "" {
		log.Fatalf("U should specify destination folder")
	}

	go findDuplicates(ch,path)
	data := mapResults(ch)
	filesToRemove := resultWorker(data, clear)

	if clear {
		removeDuplicates(filesToRemove)
	}
	fmt.Println("work complete.")
}

// Using WalkDir, introduced in Go 1.16,
// which avoids calling os.Lstat on every visited file or directory.
func findDuplicates(ch chan fileStat, path string) {
	wg := &sync.WaitGroup{}
	err := filepath.WalkDir(path, CustomWalkDirFunc(ch, wg))
	if err != nil {
		log.Println("error while walk the directory ", err)
	}
	wg.Wait()
	close(ch)
}

func CustomWalkDirFunc(ch chan fileStat, wg *sync.WaitGroup) func(string, os.DirEntry, error) error {
		return func(path string, entry os.DirEntry, err error) error {
			info, ierr := entry.Info()
			if ierr != nil {
				log.Printf("error while getting file info: %s\n", err)
			}
			// 1) т.к. ищем дубликаты не по имени, а по содержанию,
			// то ищем только файлы ненулевого размера (файлы 0 size по умолчанию идентичные -> пропускаем)
			// 2) ModeType: это "Mask for the type bits" ->
			// ModeDir|ModeSymlink|ModeNamedPipe|ModeSocket|ModeDevice|ModeCharDevice|ModeIrregular
			// т.о. если нет такого бита - то это файл, что нам и нужно
			if err == nil && info.Size() > 0 && (entry.Type() & os.ModeType == 0) {
				wg.Add(1)
				go processFile(path, info, ch, wg)
			}
			return nil // не возвращаем ошибку, возвращаем nil
		}
}

func processFile(file string, info os.FileInfo, ch chan fileStat, wg *sync.WaitGroup) {
	defer wg.Done()
	f, err := os.Open(file)
	if err != nil {
		log.Printf("cannot open file", err)
		return
	}
	defer f.Close() // намеренно для простоты, для пром поправить!

	hash := crc32.NewIEEE()
	size, err := io.Copy(hash, f)

	if size != info.Size() {
		log.Println("cannot read whole file", file)
		return
	}
	if err != nil {
		log.Println(err)
		return
	}

	ch <- fileStat{hash.Sum([]byte("mock")),info.Size(), file}
}

func mapResults(ch <-chan fileStat) map[string]*paths {
	mm := make(map[string]*paths)
	for msg := range ch {
		key:= fmt.Sprintf("%x%x", msg.size, msg.hash)
		val, ok := mm[key]
		if !ok {
			val = &paths{}
			mm[key] = val
		}
		*val = append((*val), msg.path)
	}
	return mm
}

func resultWorker(mm map[string]*paths, clear bool) (toRemove []string) {
	for _, val := range mm {
		if len(*val) > 1 {
			fmt.Printf("# number of duplicates - %d, see the list below:\n", len(*val))
			sort.Slice(*val, func(i, j int) bool {  // сортируем по короткому пути
				f1,_ := os.Stat((*val)[i]); f2,_:=os.Stat((*val)[j])
				ctf1 := f1.Sys().(*syscall.Stat_t); ctf2 := f2.Sys().(*syscall.Stat_t)
				fmt.Printf("f1 %s - %s - %d\n", (*val)[i],timespecToTime(ctf1.Ctimespec), timespecToTime(ctf1.Ctimespec).Unix());
				fmt.Printf("f2 %s - %s - %d\n", (*val)[j],timespecToTime(ctf2.Ctimespec), timespecToTime(ctf2.Ctimespec).Unix())
				return timespecToTime(ctf1.Ctimespec).Unix() < timespecToTime(ctf2.Ctimespec).Unix()
				//return len((*val)[i]) < len((*val)[j])
			})
			fmt.Printf("\t%s\n", (*val)[0])
			for _, file := range (*val)[1:] {
				fmt.Printf("\t%s\n", file)
				if clear {
					toRemove = append(toRemove,file)
				}
			}
		}
	}
	return
}

func removeDuplicates(files []string) {
	if len(files) > 1 {
		fmt.Println("\nclear flag was set to true, duplicates will be removed.\n")
	}
	wg := sync.WaitGroup{}
	var mu sync.Mutex
	wg.Add(len(files))
	for _, file := range files {
		go func(f string) {
			defer wg.Done()
			mu.Lock()
			err := os.Remove(f)
			if err !=nil {
				fmt.Printf("Cannot remove file %s. Skip.\n",f)
			} else {
				fmt.Printf("File deleted: %s.\n",f)
			}
			mu.Unlock()
		}(file)
	}
	wg.Wait()
}

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}