package search

// BinarySearch is a search algorithm with time complexity of O(log(n))
func BinarySearch(arr []int, sought int) int {
	var (
		low, mid, guess int
	)
	high := len(arr) - 1

	for low < high {
		mid = (low + high) / 2
		guess = arr[mid]
		if guess == sought {
			return mid
		}
		if guess < sought {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
