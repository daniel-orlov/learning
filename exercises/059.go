package main

//create a func with the identifier foo that
//takes in a variadic parameter of type int
//pass in a value of type []int into your func (unfurl the []int)
//returns the sum of all values of type int passed in
//create a func with the identifier bar that
//takes in a parameter of type []int
//returns the sum of all values of type int passed in

import "fmt"

func main() {
	nums := []int{
		3, 1, 4, 1, 5, 9, 2, 6, 5, 7, 5, 8, 9, 3, 9, 6, 5, 4, 8, 2, 1, 5,
	}
	sum := foo(nums...)
	fmt.Println(sum)
	sum1 := bar(nums)
	fmt.Println(sum1)
}

func foo(nums ...int) int {
	total := 0
	for _, i := range nums {
		total += i
	}
	return total
}

func bar(nums []int) int {
	total := 0
	for _, i := range nums {
		total += i
	}
	return total
}
