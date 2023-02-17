package main

import "fmt"

//Continue using slices (deleting from a slice (using slicing and append()))

func main() {
	pi := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 3, 9}
	//deleting the only 4
	pi = append(pi[:2], pi[3:]...)
	fmt.Println(pi) //that's NSFW
}
