package main

import "sync"

func main() {
	ch:=make(chan struct{})

	wg:=sync.WaitGroup{}
	wg.Add(2)
	go func() {
		_ = <-ch
		wg.Done()
	}()

	go func() {
		ch <- struct{}{}
		wg.Done()
	}()
	wg.Wait()
}
