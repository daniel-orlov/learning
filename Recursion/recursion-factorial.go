package recursion

import "fmt"

func factorial(x int, prefix string) int {
	fmt.Println(prefix, "THE VALUE OF X IS CURRENTLY: ", x)
	if x == 1 {
		fmt.Println(prefix, "Reached the base case")
		return x
	}
	prefix += "|\t"
	return x * factorial(x-1, prefix)
}
