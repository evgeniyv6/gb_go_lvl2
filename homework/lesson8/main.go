package main

import (
	"flag"
	"fmt"
	"github.com/evgeniyv6/homework/lesson8/duplicates"
	"log"
	"os"
	"path/filepath"
)

const count = 1000

func main() {
	var (
		ch = make(chan duplicates.FileStat, count)
		path string
		clear bool
	)

	log.SetFlags(log.LUTC | log.Lmicroseconds | log.Lshortfile | log.Ldate)

	flag.Usage = func() {
		_, err := fmt.Fprintf(os.Stderr, "Usage of %s: This program searches for duplicate files in a folder and subfolders " +
			"and, if necessary, deletes them\n", filepath.Base(os.Args[0]))
		if err != nil {
			fmt.Println("error:", err)
		}
		flag.PrintDefaults()
	}
	flag.StringVar(&path,"folder", "", "required parameter. path to a folder.")
	flag.BoolVar(&clear,"clear", false,"bool parameter, default = false. if set to true - duplicates will be removed.")
	flag.Parse()

	if path == "" {
		fmt.Println("U should specify destination folder")
		os.Exit(1)
	}

	go duplicates.FindDuplicates(ch,path, duplicates.FS)
	data := duplicates.MapResults(ch)
	filesToRemove := duplicates.ResultWorker(data, clear)

	if clear {
		duplicates.RemoveDuplicates(filesToRemove)
	}
	fmt.Println("work complete.")
}
