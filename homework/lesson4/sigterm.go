package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const timeout time.Duration = 2

func SigtermExtTimeout() {
	cancelSig := make(chan os.Signal)
	workers := make(chan struct{})
	done := make(chan bool,1)

	signal.Notify(cancelSig, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		for i:=0; i<100;i++ {
			workers <- struct{}{}
			time.Sleep(1*time.Second)
		}
	}()

	go func() {
		for {
			fmt.Println("get from worker", <- workers)
		}

	}()

	go func() {
		sig := <-cancelSig
		fmt.Printf("caught %v signal\n",sig)
		time.Sleep(timeout*time.Second)
		done <- true
	}()
	fmt.Printf("graceful exit - %v, after %d sec", <-done, timeout)

	time.Sleep(25*time.Second)
}


func SigtermExtTimeout2() {
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

func main() {
	SigtermExtTimeout2()
}
