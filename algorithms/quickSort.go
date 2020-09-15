package main

import (
	"fmt"
	"math/rand"
)

func QuickSort(arr []int) []int {
	//base case
	if len(arr) <= 1 {
		return arr
	}
	//picking a median randomly
	median := arr[rand.Intn(len(arr))]
	//allocating 3 slices
	lowPart := make([]int, 0, len(arr))
	middlePart := make([]int, 0, len(arr))
	highPart := make([]int, 0, len(arr))
	//sorting
	for _, item := range arr{
		switch{
		case item < median:
			lowPart = append(lowPart, item)
		case item == median:
			middlePart = append(middlePart, item)
		case item > median:
			highPart = append(highPart, item)
		}
	}
	lowPart = QuickSort(lowPart)
	highPart = QuickSort(highPart)
	lowPart = append(lowPart, middlePart...)
	lowPart = append(lowPart, highPart...)
	return lowPart
}

func RandomArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i <= n-1; i++ {
		arr[i] = rand.Intn(n)
	}
	return arr
}

func main() {
	arr := RandomArray(30)
	fmt.Println("Initial array:", arr)
	fmt.Println("Sorted  array:", QuickSort(arr))
}
