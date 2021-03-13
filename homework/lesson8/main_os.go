package main

import (
	"fmt"
	"os"
)

func main() {
	//dir, err := os.Getwd()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(dir)
	//
	//err =os.Mkdir("tempGo", os.ModePerm)
	//if err != nil {
	//	panic(err)
	//}

	err := os.Chdir("./tempGo")
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fmt.Println(dir)


	f1, err := os.Create("file1")
	if err != nil {
		panic(err)
	}

	_, err = f1.WriteString("semoe data")

	if err != nil {
		panic(err)
	}

	_ = f1.Close()


	f2, err := os.Open("file1")
	if err != nil {}
	var data = make([]byte, 9)

	_, err = f2.Read(data)
	if err!= nil {}

	fmt.Println(string(data))



}
