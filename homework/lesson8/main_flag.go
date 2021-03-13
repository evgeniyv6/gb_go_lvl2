package main

import (
	"flag"
	"fmt"
)

var (
	a *bool
	b *int
	c *string
)

// go run main_flag.go -a=false
func init() {
	a = flag.Bool("a", true, "bool param")
	b = flag.Int("b", 5, "int param")
	c = flag.String("c", "empty", "str param")
	flag.Parse()
}

func main() {
	fmt.Println(*a, *b, *c)
}
