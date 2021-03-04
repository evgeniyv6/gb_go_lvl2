package main

import (
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
}

func main() {
	SigtermExtTimeout()
}
