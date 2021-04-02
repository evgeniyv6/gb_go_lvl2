// Package for duplicates searches and removes duplicate files (sorted by ctime)
// Using WalkDir, introduced in Go 1.16,
// which avoids calling os.Lstat on every visited file or directory.
//
// FindDuplicates(ch chan FileStat, path string)
// MapResults(ch <-chan FileStat) map[string]*Paths
// ResultWorker(mm map[string]*Paths, clear bool) (toRemove []string)
// RemoveDuplicates(files []string)

package duplicates

import (
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

type Paths []string

type FileStat struct {
	Hash []byte
	Size int64
	Path string
}

func FindDuplicates(ch chan FileStat, path string) {
	wg := &sync.WaitGroup{}
	err := filepath.WalkDir(path, customWalkDirFunc(ch, wg))
	if err != nil {
		log.Println("error while walk the directory ", err)
	}
	wg.Wait()
	close(ch)
}

func customWalkDirFunc(ch chan FileStat, wg *sync.WaitGroup) func(string, os.DirEntry, error) error {
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
		if err == nil && info.Size() > 0 && (entry.Type()&os.ModeType == 0) {
			wg.Add(1)
			go processFile(path, info, ch, wg)
		}
		return nil // не возвращаем ошибку, возвращаем nil
	}
}

func processFile(file string, info os.FileInfo, ch chan FileStat, wg *sync.WaitGroup) {
	defer wg.Done()
	f, err := os.Open(file)
	if err != nil {
		log.Printf("cannot open file %s", err)
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

	ch <- FileStat{hash.Sum(nil), info.Size(), file}
}

func MapResults(ch <-chan FileStat) map[string]*Paths {
	mm := make(map[string]*Paths)
	format := fmt.Sprintf("%%016X:%%%dX", crc32.Size*2) // == "%016X:%40X"
	for msg := range ch {
		key := fmt.Sprintf(format, msg.Size, msg.Hash)
		val, ok := mm[key]
		if !ok {
			val = &Paths{}
			mm[key] = val
		}
		*val = append((*val), msg.Path)
	}
	return mm
}

func ResultWorker(mm map[string]*Paths, clear bool) (toRemove []string) {
	for _, val := range mm {
		if len(*val) > 1 {
			fmt.Printf("# number of duplicates - %d, see the list below:\n", len(*val))
			sort.Slice(*val, func(i, j int) bool { // сортируем по ctime
				f1, err := os.Stat((*val)[i])
				if err != nil {
					fmt.Printf("cannot get file stat %s. Skip.\n", (*val)[i])
				}
				f2, err := os.Stat((*val)[j])
				if err != nil {
					fmt.Printf("cannot get file stat %s. Skip.\n", (*val)[j])
				}
				ctimef1 := f1.Sys().(*syscall.Stat_t).Ctimespec
				ctimef2 := f2.Sys().(*syscall.Stat_t).Ctimespec
				return timespecToTime(ctimef1).Unix() < timespecToTime(ctimef2).Unix()
				// return len((*val)[i]) < len((*val)[j]) // либо сортируем по короткому пути, в зависимости от ТЗ,
				// либо добавить признак выбора сортировки для пользователя - TBD
			})
			fmt.Printf("\t%s\n", (*val)[0])
			for _, file := range (*val)[1:] {
				fmt.Printf("\t%s\n", file)
				if clear {
					toRemove = append(toRemove, file)
				}
			}
		}
	}
	return
}

func RemoveDuplicates(files []string) {
	if len(files) > 1 {
		fmt.Printf("\nclear flag was set to true, duplicates will be removed.\n")
	}
	wg := sync.WaitGroup{}
	var mu sync.Mutex
	wg.Add(len(files))
	for _, file := range files {
		go func(f string) {
			defer wg.Done()
			mu.Lock()
			err := os.Remove(f)
			if err != nil {
				fmt.Printf("Cannot remove file %s. Skip.\n", f)
			} else {
				fmt.Printf("File deleted: %s.\n", f)
			}
			mu.Unlock()
		}(file)
	}
	wg.Wait()
}

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix((ts.Sec), (ts.Nsec))
}
