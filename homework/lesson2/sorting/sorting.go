// Package sorting implements functions to demonstrate godoc features
//
// The ShakerSortingMemPos implementation of the shaker sorting algorithm
//
// ShakerSortingMemPos () int
package sorting

// ShakerSortingMemPos returns the number of steps required to sort a slice using a shaker sort
func ShakerSortingMemPos(sl *[]int) (counter int) {
	var swapped = true
	var startIndx int; var startSwap int
	var endIndx, endSwap = len(*sl) - 1, len(*sl) - 1

	for swapped {
		swapped = false
		for i:=startIndx; i < endIndx; i++ {
			counter++
			if (*sl)[i] > (*sl)[i+1] {
				counter++
				(*sl)[i], (*sl)[i+1] = (*sl)[i+1], (*sl)[i]
				swapped = true
				endSwap = i
			}
		}
		if !swapped {
			break
		} else {
			swapped = false
			endIndx = endSwap
		}

		for i:= endIndx; i > startIndx;i-- {
			counter++
			if (*sl)[i] < (*sl)[i-1] {
				counter++
				(*sl)[i], (*sl)[i-1] = (*sl)[i-1], (*sl)[i]
				swapped = true
				startSwap = i
			}
		}
		startIndx = startSwap
	}
	return
}
