package main

import "fmt"

func main() {
	const literal = "🍎🍎🍎\uF8FF\uF8FF"

	fmt.Printf("%s\n", literal)
	fmt.Printf("%q\n", literal)
	fmt.Printf("%x\n", literal)

	for i:=0; i<len(literal); i++ {
		fmt.Printf("%U ", literal[i])
	}
	fmt.Println()
	for _,c:=range literal {
		fmt.Printf("%U ",c)
	}
	fmt.Println()
}
