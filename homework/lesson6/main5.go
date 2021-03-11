package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime/trace"
	"sync"
)

// to check trace

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	var (
		ch = make(chan int)
		done = make(chan struct{})
		wg = sync.WaitGroup{}
	)
	wg.Add(1)
	go func() {
		defer close(done)
		defer wg.Done()
		for i:=0; i<10; i++ {
			ch<-rand.Int()
		}
	}()

	for i:=0;i<3;i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case v:=<-ch:
					fmt.Println(i,v)
				case <-done:
					return
				}
			}
		}(i)
	}
	wg.Wait()
}
