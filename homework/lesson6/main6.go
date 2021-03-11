package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

// race

const count = 1000

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	var (
		counter int
		wg sync.WaitGroup
		lock sync.Mutex
	)
	wg.Add(count)

	for i:=0; i<count;i++ {
		go func() {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			counter ++
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
