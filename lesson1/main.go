package main

import (
	"errors"
	"fmt"
	"runtime/debug"
	"time"
)

type ErrWithStack struct {
	txt string
	stackTrace string
}

func (e *ErrWithStack) Error() string {
	return fmt.Sprintf("err: %s\n stack trace: %s", e.txt, e.stackTrace)
}

func New(text string) error {
	return &ErrWithStack{
		txt: text,
		stackTrace: string(debug.Stack()),
	}
}

func main() {
	fmt.Println("hi from lesson1 Golvl2")
	example4()
}


// example 1
func example1() {
	var err error
	err = errors.New("Ex1: my err1")
	fmt.Println(err)

	err = New("Ex1: my err2")
	fmt.Println(err)

	err = fmt.Errorf("foo func %w", foo())
	fmt.Println(err)
}

func foo() error {
	return fmt.Errorf("found err: %w", errors.New("some err text"))
}



// example 2
var findErr = errors.New("investigate err")

func example2() {
	err := fmt.Errorf("exmpl2: %w", findErr)
	fmt.Println(errors.Is(err, findErr))
}

// example3()
func example3() {
	err1 := fmt.Errorf("found err: %w", New("my err!!!"))
	err2:= &ErrWithStack{}
	if ok:= errors.As(err1, &err2); ok {
		fmt.Println("it is err with stack trace", err2)
	} else {
		fmt.Println("fail")
	}
}


// example4
func example4() {
	//panic(errors.New("some panic err"))
	bar()
	fmt.Println("After FOO")
}

func bar() {
	defer func() {
		if v:=recover(); v !=nil {
			fmt.Println(v)
		}
	}()

	go func() {
		var a int
		fmt.Println(1/a)
	}()

	time.Sleep(3*time.Second)

}
