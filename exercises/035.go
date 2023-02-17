package main

//Using a COMPOSITE LITERAL:
//create a SLICE of TYPE int
//assign 10 VALUES
//Range over the slice and print the index-value pairs out.
//Using format printing
//print out the TYPE of the slice

import "fmt"

func main() {
	the_slice := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}

	for i, v := range the_slice {
		fmt.Printf("Index:\t%v\t|\tValue:\t%v\n", i, v)
	}
	fmt.Printf("Type of the slice is %T", the_slice)
}
