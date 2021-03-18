package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	a *bool
	b *int
	c *string
)

// go run main_flag.go -a=false
func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "test msg\n")
		flag.PrintDefaults()
	}
	a = flag.Bool("a", true, "bool param")
	b = flag.Int("b", 5, "int param")
	c = flag.String("c", "empty", "str param")
	flag.Parse()
}

func main() {
	fmt.Println(*a, *b, *c)
}
