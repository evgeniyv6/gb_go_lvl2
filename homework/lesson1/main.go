package main

import (
	"fmt"
	"github.com/evgeniyv6/homework/lesson1/fileclose"
	"github.com/evgeniyv6/homework/lesson1/gopanic"
)

func main() {
	fmt.Println("Hello from homework 1.")

	// exercise 1
	gopanic.GoroutinePanicCatcher()

	if err := fileclose.FileClose("testfile.txt"); err != nil {
		fmt.Println("Cannot work with file.")
	}

}
