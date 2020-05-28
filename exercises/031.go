package main

import "fmt"

//Continue using slices (creating a slice of a specific size using make())

func main() {
	week := make([]int, 7, 31)
	fmt.Println(week)
	fmt.Println(len(week))
	fmt.Println(cap(week))
}
