package recursion

import "fmt"

func Sum(arr []int, prefix string) int {
	fmt.Println(prefix, "ARRAY IS CURRENTLY: ", arr)
	if len(arr) == 1 {
		fmt.Println(prefix, "Reached the base case")
		return arr[0]
	}
	prefix += "|\t"
	return arr[0] + Sum(arr[1:], prefix)
}
