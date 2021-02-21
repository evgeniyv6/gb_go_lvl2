package sorting

import "fmt"

func Example() {
	sl := []int{8,7,3,2,1,2,4,5}
	ShakerSortingMemPos(&sl)
	fmt.Println(sl)
}
