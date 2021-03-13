package main

import (
	"fmt"
	"sort"
)

type User struct {
	Id string
	Name string
	Age int
}

type SortedUsers []User

func (u SortedUsers) Len() int {
	return len(u)
}

func (u SortedUsers) Less(i,j int) bool {
	return (u)[i].Age < (u)[j].Age
}

func (u SortedUsers) Swap(i,j int) {
	(u)[i], (u)[j] =(u)[j], (u)[i]
}



func main() {
	ff:= []float64{1.,.4,2.5}
	sort.Float64s(ff)
	fmt.Println(ff)

	ii:= []int{4,3,5,6,2,5}
	sort.Ints(ii)
	fmt.Println(ii)

	ss:= []string{"1", "a", "s", "3"}
	sort.Strings(ss)
	fmt.Println(ss)

	uu:= []User{
		{"1", "Alex",25},
		{"2", "John",11},
		{"3", "Ann",39},
	}

	sort.Sort(SortedUsers(uu))

	fmt.Println(uu)


}
