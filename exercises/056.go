package main

//Using callback, passing a function as an argument to a function

import "fmt"

func main() {
	boxes := []int {
		1,31,54,54,12,54,56,5,87,32,
	}
	resultAllNums := sum(boxes...)
	fmt.Println(resultAllNums)
	resultEvenNums := sum((evenOrOdd(true, boxes...))...)
	fmt.Println(resultEvenNums)
	resultOddNums := sum((evenOrOdd(false, boxes...))...)
	fmt.Println(resultOddNums)
}

func sum(nums... int) int {
	total := 0
	for _, v := range nums{
		total += v
	}
	return total
}

func evenOrOdd(modeEven bool, nums... int) []int {
	var result []int
	if modeEven {
		for _, num := range nums{
			if num % 2 == 0 {
				result = append(result, num)
			}
		}
	}else{
		for _, num := range nums{
			if num % 2 == 1 {
				result = append(result, num)
			}
		}
	}
	return result
}