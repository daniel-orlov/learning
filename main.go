package main

import (
	"fmt"
)

func main() {
	//Binary Search
	binArr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sought := 5
	index := BinarySearch(binArr, sought)
	fmt.Println("BINARY SEARCH:\nArray: ", binArr)
	fmt.Printf("Value '%v' found in arr at index %v\n", sought, index)

	//SelectionSort
	selSortArr := []int{3, 1, 4, 5, 9, 2, 6, 8, 7, 0}
	fmt.Println("\nSELECTION SORT:\nArray: ", selSortArr)
	sortedArr := SelectionSort(selSortArr)
	fmt.Println("Sorted array: ", sortedArr)
}
