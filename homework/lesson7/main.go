package main

import (
	"fmt"
	"strconv"
)

type Walker interface {
	Walk()
}

type Speaker interface {
	Hello()
}

type Man struct {
	Greeting string
	Meters int
}

func (h *Man) Hello() {
	fmt.Println(h.Greeting)
}

func (h *Man) Walker() {
	fmt.Println(h.Meters)
}

func (h *Man) String() string  {
	return fmt.Sprintf("Greet %s %d", h.Greeting, h.Meters)
}

func main() {
	fmt.Println(" hi from l7")

	h:= Man{"Salut", 10,}
	s:=Speaker(&h)

	s.Hello()

	//w, ok := s.(Walker)
	//if !ok {
	//	fmt.Printf("t a %T\n", s)
	//	return
	//}
	//w.Walk()

	fmt.Println(s)
	fmt.Println(ToString(10))
	fmt.Println(ToString("22"))
	fmt.Println(ToString(h))
	fmt.Println(ToString(3.12))

}

func ToString(in interface{}) string {
	switch v:= in.(type) {
	case int:
		return strconv.Itoa(v)
	case string:
		return v
	case Man:
		return v.String()
	default:
		return "???"
	}

}

