package main

import (
	"fmt"
	"reflect"
)

type MyFloat float64

func main()  {
	fmt.Println("reflect")

	var x MyFloat = 3.1415
	valOf := reflect.ValueOf(x)
	fmt.Println(valOf)
	fmt.Println(valOf.Type())
	fmt.Println(valOf.Kind())

	fmt.Printf("%v, %T\n", valOf, valOf)

	i:= valOf.Interface()
	y, ok := i.(MyFloat)
	if !ok {
		fmt.Println("ta err")
		return
	}
	fmt.Println(y)
	fmt.Printf("%v, %T\n",y, y)

//////
	var xx float64 = 3.1
	fmt.Println(xx)
	valOf2:= reflect.ValueOf(&xx)
	fmt.Println(valOf2)
	v2:=valOf2.Elem() // разыменовывает
	fmt.Println(v2.CanSet())
	v2.SetFloat(2.12)
	fmt.Println(v2)

	fmt.Println(xx)



}
