package main

//Three different ways to count 5!

import "fmt"

func main() {
	//Using a recursive function factorial
	fmt.Println(factorial(5))
	//Using a function with a for-loop
	fmt.Println(forLoopFactorial(5))
	//Hard-coding it like there's no tomorrow
	fmt.Println(5*4*3*2*1)
}

func forLoopFactorial(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}

func factorial(i int) int {
	if i == 0 {
		return 1
	}
	return i * factorial(i-1)
}
