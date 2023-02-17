package main

//Using a COMPOSITE LITERAL:
//create an ARRAY which holds 5 VALUES of TYPE int
//assign VALUES to each index position.
//Range over the array and print the index-value pairs out.
//Using format printing
//print out the TYPE of the values too

import "fmt"

func main() {
	the_array := [5]int{4, 3, 2, 1, 0}
	//fmt.Println(the_array)
	for i, v := range the_array {
		fmt.Printf("Index:\t%v\t|\tValue:\t%v\t\n", i, v)
	}
	fmt.Printf("Type of the array is %T", the_array)
}
