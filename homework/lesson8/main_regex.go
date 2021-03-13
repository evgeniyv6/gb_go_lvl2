package main

import (
	"fmt"
	"regexp"
)

func main() {
	var (
		s ="jh3g24kj2hbdk2j3hbd2k3jh2332kj3hb2kjh3b"
		re = regexp.MustCompile(`([0-9])+`)
	)

	ss := re.FindAllString(s, -1)

	for _, s:=range ss {
		fmt.Println(s)
	}
}
