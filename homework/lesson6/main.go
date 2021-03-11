package main

import (
	"fmt"
	"github.com/evgeniyv6/homework/lesson6/iqueue"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

const count = 10

func exercise1() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	var (
		q iqueue.Queue
		wg sync.WaitGroup
	)
	wg.Add(count)
	for i:=0;i<count;i++ {
		go func(i int) {
			defer wg.Done()
			q.Enqueue(i)
		}(i)
	}

	wg.Wait()
	q.Sort()
	fmt.Println(q.Items)
}

func exercise2() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	for i:=0;i<1000* count;i++{
		go func() {
			time.Sleep(1*time.Minute)
		}()
		if i%10 == 0 {
			runtime.Gosched()
		}
	}
}

func main() {
	// exercise1()
	exercise2()
}
