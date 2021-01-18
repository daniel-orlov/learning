package sorting

import (
	"math/rand"
)

//QuickSort implementation with recursion
//TBH, looks a bit like memalloc hell to me
func QuickSort(arr []int) []int {
	//base case
	if len(arr) < 2 {
		return arr
	}
	pivotIndex := rand.Intn(len(arr))
	pivot := arr[pivotIndex]
	less := make([]int, 0, len(arr)-1)
	more := make([]int, 0, len(arr)-1)

	for _, i := range arr {
		if i < pivot {
			less = append(less, i)
		} else if i > pivot {
			more = append(more, i)
		}
	}
	return append(append(QuickSort(less), pivot), QuickSort(more)...)
}

//TODO: implement in place QuickSort
