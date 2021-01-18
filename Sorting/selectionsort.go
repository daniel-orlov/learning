package sorting

/*Time Complexity: O(n^2)
 */
func SelectionSort(arr []int) []int {
	sortedArr := make([]int, 0, len(arr))
	minIndex := 0
	for range arr {
		minIndex = findIndexMin(arr)
		sortedArr = append(sortedArr, arr[minIndex])
		arr = PopItem(arr, minIndex)
	}
	return sortedArr
}

func findIndexMin(arr []int) int {
	length := len(arr)
	//handling base cases
	minIndex := 0
	min := arr[minIndex]
	if length == 1 {
		return minIndex
	}
	if length == 2 {
		if min < arr[1] {
			return min
		}
		return 1
	}
	//iterating and finding a minimal value and its index
	for ind, val := range arr {
		if val < min {
			minIndex = ind
			min = val
		}
	}

	return minIndex
}

func PopItem(arr []int, index int) []int {
	newArr := make([]int, 0, len(arr)-1)
	newArr = append(arr[:index], arr[index+1:]...)
	return newArr
}
