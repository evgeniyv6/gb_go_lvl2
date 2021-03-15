//go:generate ./genstack MyInter
package main

import "fmt"
type MyInter interface {}

func main() {
	var o, t, th MyInter = 1,2,3
	st := NewMyInterStack()
	st.Push(o); st.Push(t); st.Push(th)
	st.Pop()

	fmt.Println(st)

}
