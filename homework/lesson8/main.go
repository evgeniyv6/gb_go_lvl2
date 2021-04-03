package main

import (
	"flag"
	"fmt"
	"github.com/evgeniyv6/homework/lesson8/duplicates"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"sync"
)

const count = 1000

func main() {
	var (
		ch    = make(chan duplicates.FileStat, count)
		errCh = make(chan error)
		wg    = sync.WaitGroup{}
		path  string
		clear bool
	)

	// log.SetFlags(log.LUTC | log.Lmicroseconds | log.Lshortfile | log.Ldate)
	logrus.SetFormatter(&logrus.TextFormatter{})
	standartFields := logrus.Fields{"name": "find files duplicates"}
	hlog := logrus.WithFields(standartFields)

	flag.Usage = func() {
		_, err := fmt.Fprintf(os.Stderr, "Usage of %s: This program searches for duplicate files in a folder and subfolders "+
			"and, if necessary, deletes them\n", filepath.Base(os.Args[0]))
		if err != nil {
			fmt.Println("error:", err)
		}
		flag.PrintDefaults()
	}
	flag.StringVar(&path, "folder", "", "required parameter. path to a folder.")
	flag.BoolVar(&clear, "clear", false, "bool parameter, default = false. if set to true - duplicates will be removed.")
	flag.Parse()

	if path == "" {
		fmt.Println("U should specify destination folder")
		os.Exit(1)
	}

	go func() {

		for err := range errCh {
			wg.Add(1)
			if err != nil {
				hlog.Errorf("CATCH ERROR %v", err)
			}
			wg.Done()
		}

	}()

	go duplicates.FindDuplicates(ch, path, errCh)
	data := duplicates.MapResults(ch)
	filesToRemove := duplicates.ResultWorker(data, clear, errCh)

	if clear {
		duplicates.RemoveDuplicates(filesToRemove, errCh)
	}
	wg.Wait()
	logrus.Info("work complete.")
}
