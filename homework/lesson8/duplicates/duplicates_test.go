package duplicates

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"hash/crc32"
	"io"
	"os"
	"reflect"
	"sort"
	"sync"
	"syscall"
	"testing"
	"time"
)

const testFolder = "/tmp/duplicatesGotest/"

// go test -v .
func TestFindDuplicates(t *testing.T) {
	err := os.RemoveAll(testFolder)
	if err != nil {
		panic(err)
	}

	mm := prepareDataMap()
	ch := make(chan FileStat, 10)
	wg := sync.WaitGroup{}
	errCh := make(chan error)

	logrus.SetFormatter(&logrus.TextFormatter{})
	standartFields := logrus.Fields{"name": "find files duplicates"}
	hlog := logrus.WithFields(standartFields)

	go func() {
		for err := range errCh {
			wg.Add(1)
			if err != nil {
				hlog.Errorf("CATCH ERROR %v", err)
			}
			wg.Done()
		}

	}()

	go FindDuplicates(ch, testFolder, errCh)
	data := MapResults(ch)

	fmt.Println("Test data:")
	_ = ResultWorker(mm, false, errCh)
	fmt.Println("Func data:")
	_ = ResultWorker(data, false, errCh)

	mm1 := make(map[string]Paths, len(mm))
	mm2 := make(map[string]Paths, len(mm))

	for k, v := range mm {
		sort.Slice(*v, func(i, j int) bool { // сортируем по ctime
			f1, err := os.Stat((*v)[i])
			if err != nil {
				fmt.Printf("cannot get file stat %s. Skip.\n", (*v)[i])
			}
			f2, err := os.Stat((*v)[j])
			if err != nil {
				fmt.Printf("cannot get file stat %s. Skip.\n", (*v)[j])
			}
			ctimef1 := f1.Sys().(*syscall.Stat_t).Ctimespec
			ctimef2 := f2.Sys().(*syscall.Stat_t).Ctimespec
			return timespecToTime(ctimef1).Unix() < timespecToTime(ctimef2).Unix()
		})
		mm1[k] = *v
	}
	for k, v := range data {
		mm2[k] = *v
	}

	eq := reflect.DeepEqual(mm1, mm2)

	if eq {
		fmt.Println(mm1)
		fmt.Println(mm2)
		fmt.Println("Maps're equal.")
	} else {
		t.Errorf("Test failed maps not equal 1 - %v\n 2 - %v\n", mm1, mm2)
	}
	wg.Wait()
}

func prepareDataMap() map[string]*Paths {
	err := os.MkdirAll(testFolder, 0755)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(testFolder+"inner", 0755)
	if err != nil {
		panic(err)
	}

	var files = []string{testFolder + "file1.txt",
		testFolder + "file2.txt", testFolder + "inner/file3.txt"}

	for _, f := range files {
		fileCreator(f, "duplicate")
		time.Sleep(10 * time.Millisecond)
	}
	fileCreator(files[len(files)-1], "non duplicate")

	mm := make(map[string]*Paths)
	format := fmt.Sprintf("%%016X:%%%dX", crc32.Size*2) // == "%016X:%40X"
	for _, f := range files {
		hash := crc32.NewIEEE()
		fo, err := os.Open(f)
		if err != nil {
			panic(err)
		}
		_, err = io.Copy(hash, fo)
		if err != nil {
			panic(err)
		}
		fi, err := os.Stat(f)
		if err != nil {
			panic(err)
		}
		key := fmt.Sprintf(format, fi.Size(), hash.Sum(nil))

		val, ok := mm[key]
		if !ok {
			val = &Paths{}
			mm[key] = val
		}

		*val = append((*val), f)

	}
	return mm
}

func fileCreator(path string, text string) {
	f1, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	_, err = f1.WriteString(text)
	if err != nil {
		panic(err)
	}
	err = f1.Close()
	if err != nil {
		panic(err)
	}
}
