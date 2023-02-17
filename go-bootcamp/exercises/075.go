package main

/*
Sorting slices using pkg sort
*/

import (
	"fmt"
	"sort"
)

func main() {
	ss := []string{
		"John", "Paul", "George", "Ringo",
	}
	si := []int{
		3, 7, 2, 4, 1, 5, 8, 6,
	}
	fmt.Println("_________UNSORTED_________")
	fmt.Println(ss)
	fmt.Println(si)
	fmt.Println("__________SORTED__________")
	sort.Strings(ss)
	sort.Ints(si)
	fmt.Println(ss)
	fmt.Println(si)

}
