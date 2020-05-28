package main

import "fmt"

func merge(a []int, b []int) []int  {
	var res = make([]int, len(a) + len(b))
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			res[i+j] = a[i]
			i++
		} else {
			res[i+j] = b[j]
			j++
		}
	}
	for i < len(a) {
		res[i+j] = a[i]
		i++
	}
	for j < len(b) {
		res[i+j] = b[j]
		j++
	}
	return res
}

func Mergesort(sl []int) []int {

	if len(sl) < 2 {
		return sl
	}
	mid := len(sl) / 2
	a := Mergesort(sl[:mid])
	b := Mergesort(sl[mid:])
	return merge(a, b)
}

func main()  {
	s := []int {
		10,43,1,33,2,7,9,4,
	}
	fmt.Print(Mergesort(s))
}