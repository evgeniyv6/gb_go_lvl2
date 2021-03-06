package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
)

type s1 struct {
	a int
	b float64
	c *s2
}

type s2 struct {
	a int64
	b float32
}


func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	var (
		a *s1
		b *s2
	)

	func() {
		a = &s1{c:&s2{}}
	}()
	func() {
		b=&s2{}
	}()

	runtime.GC()
	fmt.Println(a)
	fmt.Println(b)
}