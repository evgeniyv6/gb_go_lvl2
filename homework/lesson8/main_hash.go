package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	var (
		h1 = crc32.NewIEEE()
		h2 = crc32.NewIEEE()
	)

	_, err := h1.Write([]byte("test1"))

	if err!= nil {
		panic(err)
	}

	_, err = h2.Write([]byte("test2"))

	if err!= nil {
		panic(err)
	}

	fmt.Println(h1.Sum32())
	fmt.Println(h2.Sum32())
}