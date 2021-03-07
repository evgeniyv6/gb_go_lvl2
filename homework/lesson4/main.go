package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	count = 1000
	timeout time.Duration = 2
)

func main() {
	sigtermExtTimeout()
}

func sigtermExtTimeout() {
	var (
		ctx, cancel = context.WithCancel(context.Background())
		cancelSig = make(chan os.Signal)
		workers = make(chan struct{})
		done = make(chan bool,1)

		doWork = func(num int) {
			for i := 0; i < num; i++ {
				workers <- struct{}{}
				time.Sleep(1 * time.Second)
			}
		}

		getWork = func() {
			for {
				fmt.Println("get from worker", <- workers)
			}
		}

		catchSIG = func(cancel context.CancelFunc) {
			sig := <-cancelSig
			fmt.Println("caugth sigterm", sig)
			done <- true
			cancel()
		}
	)

	signal.Notify(cancelSig, syscall.SIGTERM, syscall.SIGINT)

	go doWork(100)
	go getWork()
	go catchSIG(cancel)

	select {
	case <-ctx.Done():
		time.Sleep(timeout*time.Second)
		fmt.Println("Completed.", <-done)
		return
	}

}

func kiloGoroutins(count int) int {
	var counter int
	var workers = make (chan struct {}, 1)
	for i := 0 ; i < count ; i++ {
		workers <- struct{}{}
		go func (){
			defer func () {
				<-workers
				counter++
			}()
		}()
	}
	time.Sleep(1 * time.Second)
	return counter
}