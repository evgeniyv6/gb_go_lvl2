package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func mainOld() {
	fmt.Println("Hello from lesson4")

	var (
		ctx, cancel = context.WithCancel(context.Background())
		hr = func(cancel context.CancelFunc) {
			for t:=0; t < 4; t+=1 {
				time.Sleep(time.Second)
			}
			cancel()
			fmt.Println("hr goes home")
		}
		jobs = make(chan int)
		manager = func(ctx context.Context) {
			for job := 0; ; job+=1 {
				select {
				case <- ctx.Done():
					close(jobs)
					fmt.Println("manager goes home")
					return
				default:
					fmt.Printf("manager creates job %d\n", job)
					jobs <- job
				}
			}
		}
		resource = make(chan struct{}, 3)
		worker = func(id int) {
			defer func() { <- resource }()
			for job := range jobs {
				fmt.Printf("worker %d starts proc %d\n",id,job)
				<- time.NewTicker(time.Second).C
				fmt.Printf("worker %d complete proc %d\n",id,job)
			}
			fmt.Printf("worker %d goes home",id)
		}
	)
	go manager(ctx)
	go hr(cancel)

	for i:=0; i<cap(resource); i+=1 {
		resource<- struct{}{}
		go worker(i)
	}

	select {
	case <-ctx.Done():
		for i:=0; i<cap(resource); i+=1 {
			resource <- struct{}{}
		}
		close(resource)
		return

	}

}

