package main

import "fmt"

func main() {
	var m = make(map[int]string)

	go func() {
		m[1] = "hello"
	}()
	m[2]="world"
	for k,v :=range m {
		fmt.Println(k,v)
	}
}
