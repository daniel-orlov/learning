package recursion

import "fmt"

func Factorial(x int, prefix string) int {
	fmt.Println(prefix, "THE VALUE OF X IS CURRENTLY: ", x)
	if x == 1 {
		fmt.Println(prefix, "Reached the base case")
		return x
	}
	prefix += "|\t"
	return x * Factorial(x-1, prefix)
}
