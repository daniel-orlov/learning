package main

import "fmt"

func tailFactorial(x int, prefix string) int {
	return _tailFactorial(x-1, x, prefix)
}

func _tailFactorial(x, current int, prefix string) int {
	fmt.Println(prefix, "THE VALUE OF X IS CURRENTLY: ", x)
	fmt.Println(prefix, "PREV RESULT OF CALCULATION : ", current)
	if x == 1 {
		fmt.Println(prefix, "Reached the base case")
		return current
	}
	prefix += "|\t"
	return _tailFactorial(x-1, x*current, prefix)
}
