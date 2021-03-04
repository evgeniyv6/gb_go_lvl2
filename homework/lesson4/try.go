package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func infiniteStack() {
	for i:=0;i<2;i++{
		go func() {
			time.Sleep(5*time.Minute)
		}()
	}
	runtime.Gosched()
}

func workers() {
	var workers = make(chan struct{}, 5)
	for i:=1;i<=10;i++ {
		workers <- struct{}{}

		go func(job int) {
			defer func() {
				<- workers
			}()
			time.Sleep(1*time.Second)
			fmt.Println("job ->",job)
		}(i)
	}
}

func ticktack() {
	ticker :=  time.NewTicker(1*time.Second)
	go func() {
		for t:=range ticker.C {
			fmt.Printf("tick %v\n",t)
		}
	}()
	time.Sleep(6*time.Second)
	ticker.Stop()
}

////backup
//func kiloGoroutins() int {
//	var counter int
//	for i:=0; i < count; i++ {
//		go func() {
//			counter++
//		}()
//	}
//	return counter
//}

/// backup

const count = 1000
func kiloGoroutins() int {
	var (
		counter int
		ch = make(chan struct{})
		ctx, cancel = context.WithCancel(context.Background())
		hr = func(cancel context.CancelFunc) {
			for i:=0; i < count; i++ {
				go func() {
					ch <- struct{}{}
					counter++
				}()
			}
			cancel()
		}
	)
	hr(cancel)
	select {
	case <-ctx.Done():
		return counter
	}
}

func kiloGoroutins22() int {
		var counter int
		ch := make(chan int,1000)
		defer close(ch)
		for i:=0; i < count; i++ {
			go func() {
				ch <- i
				counter++
			}()
		}
		go func() {
			for range ch {

			}
		}()
		time.Sleep(1*time.Second)
		return counter
}

func kiloGoroutins2() int {
	var counter int
	var workers = make (chan struct {}, 1)
	for i := 0 ; i < 1000 ; i++ {
		workers <- struct {}{}
		go func (){
			fmt.Println("goro number ",i)
			defer func () {
				<-workers
				counter++
			}()
		}()
	}
	time.Sleep( 2 * time.Second)
	return counter
}


func main() {
	//ticktack()
	// infiniteStack()
	fmt.Println(kiloGoroutins2())
	// workers()


}
