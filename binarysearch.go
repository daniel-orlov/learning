package go_algorithms

func BinarySearch(list []int, sought int) int {
	var (
		low, mid, guess int
	)
	high := len(list) - 1

	for low < high {
		mid = (low + high) / 2
		guess = list[mid]
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
