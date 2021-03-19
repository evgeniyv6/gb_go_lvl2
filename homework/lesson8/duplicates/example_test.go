package duplicates

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)
const count = 1000

func Example() {
	var (
		ch = make(chan FileStat, count)
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

	go FindDuplicates(ch,path)
	data := MapResults(ch)
	filesToRemove := ResultWorker(data, clear)

	if clear {
		RemoveDuplicates(filesToRemove)
	}
	fmt.Println("work complete.")
}
