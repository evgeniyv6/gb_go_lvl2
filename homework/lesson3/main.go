package main

import (
	"fmt"
	"github.com/evgeniyv6/gbyevgenpkg"
	gbyevgenpkg2 "github.com/evgeniyv6/gbyevgenpkg/v2"
	v1 "github.com/evgeniyv6/gbyevgenpkg/slicegen"
	v2 "github.com/evgeniyv6/gbyevgenpkg/v2/slicegen"
)


func main() {
	fmt.Println(gbyevgenpkg.Version)
	newSlice := v1.GenerSlice(10)
	fmt.Println("newSlice", newSlice)
	v1.ReverseSlice(&newSlice)
	fmt.Println("Reverced slice", newSlice)

	fmt.Println(gbyevgenpkg2.Version)
	newSlice = v2.GenerSlice(10)
	fmt.Println("newSlice", newSlice)
	v2.ReverseSlice(&newSlice)
	fmt.Println("Reverced slice", newSlice)

}


