package main

import "fmt"

func main() {
	//array size is a part of its type
	//init array populated with def values
	var arr0 [3]int // [0,0,0]
	fmt.Println("array arr0 short", arr0) 		// [0 0 0]
	fmt.Printf("array arr0 short %v\n", arr0) 	// [0 0 0]
	fmt.Printf("array arr0 full %#v\n", arr0)	// [4]int{0, 0, 0}

	//can use const when declaring size, can't use var
	const size = 4
	var arr1 [size / 2]bool
	fmt.Printf("array arr1 full %#v\n", arr1)

	//declaring via literal
	arr2 := [...]int{0, 1, 1, 2, 3, 5, 8}
	fmt.Printf("array arr2 full %#v\n", arr2)
}
