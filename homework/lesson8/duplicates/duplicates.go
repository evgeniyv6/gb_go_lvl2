// Package duplicates searches and removes duplicate files
// Using WalkDir, introduced in Go 1.16,
// which avoids calling os.Lstat on every visited file or directory.
//
// FindDuplicates(ch chan FileStat, path string)
// MapResults(ch <-chan FileStat) map[string]*paths
// ResultWorker(mm map[string]*paths, clear bool) (toRemove []string)
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
)

type paths []string

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
		if err == nil && info.Size() > 0 && (entry.Type() & os.ModeType == 0) {
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

	ch <- FileStat{hash.Sum([]byte("mock")),info.Size(), file}
}

func MapResults(ch <-chan FileStat) map[string]*paths {
	mm := make(map[string]*paths)
	for msg := range ch {
		key:= fmt.Sprintf("%x%x", msg.Size, msg.Hash)
		val, ok := mm[key]
		if !ok {
			val = &paths{}
			mm[key] = val
		}
		*val = append((*val), msg.Path)
	}
	return mm
}

func ResultWorker(mm map[string]*paths, clear bool) (toRemove []string) {
	for _, val := range mm {
		if len(*val) > 1 {
			fmt.Printf("# number of duplicates - %d, see the list below:\n", len(*val))
			sort.Slice(*val, func(i, j int) bool {  // сортируем по короткому пути
				return len((*val)[i]) < len((*val)[j])
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

func RemoveDuplicates(files []string) {
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
