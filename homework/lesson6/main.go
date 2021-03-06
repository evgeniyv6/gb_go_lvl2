package main

import (
	"fmt"
	"sync"
)

func main() {
	wg:= sync.WaitGroup{

	}
	wg.Add(2)
	go func() {
		defer wg.Done()
		c:='a'
		for i:=0;i<50;i++ {
			fmt.Printf("%c", c + int32(i%26))
		}
	}()

	go func() {
		defer wg.Done()
		var n='0'
		for i:=0;i<50;i++ {
			fmt.Printf("%c", n+int32(i%10))
		}
	}()

	wg.Wait()
}
