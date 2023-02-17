package go_algorithms

import (
	"fmt"

	"recursion"
	"search"
	"sorting"
)

func demoBinarySearch() {
	binArr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sought := 5
	index := search.BinarySearch(binArr, sought)
	fmt.Println("BINARY SEARCH:\nArray: ", binArr)
	fmt.Printf("Value '%v' found in arr at index %v\n", sought, index)
}

func demoSelectionSort() {
	selSortArr := []int{3, 1, 4, 5, 9, 2, 6, 8, 7, 0}
	fmt.Println("\nSELECTION SORT:\nArray: ", selSortArr)
	sortedArr := sorting.SelectionSort(selSortArr)
	fmt.Println("Sorted array: ", sortedArr)
}

func demoRecusionFactorial() {
	fmt.Println("\nRECURSIVE FACTORIAL CALCULATION:")
	x := 7
	result := recursion.factorial(x, "|\t")
	fmt.Printf("RESULT: Factorial %v = %v\n", x, result)
}

func demoRecursionSumArr() {
	fmt.Println("\nRECURSIVE SUM CALCULATION:")
	arr := []int{5, 7, 9, 12, 21}
	result := recursion.Sum(arr, "|\t")
	fmt.Printf("RESULT: Sum of elements in an array %v = %v\n", arr, result)
}

func demoTailRecursionFactorial() {
	fmt.Println("\nTAIL-RECURSIVE FACTORIAL CALCULATION:")
	x := 7
	result := tailFactorial(x, "|\t")
	fmt.Printf("RESULT: Factorial %v = %v\n", x, result)
}

func demoQuickSort() {
	quickSortArr := []int{3, 1, 4, 5, 9, 2, 6, 8, 7, 0}
	fmt.Println("\nQUICK SORT:\nArray: ", quickSortArr)
	quickSortedArr := QuickSort(quickSortArr)
	fmt.Println("Sorted array: ", quickSortedArr)
}
