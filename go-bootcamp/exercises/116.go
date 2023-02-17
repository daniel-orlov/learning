package main

/*
Simple multiplication function – some future testing material
*/

import "fmt"

func main() {
	r1 := multiply(1, 4, 8, 6)
	r2 := multiply(645, 897, 385, -66)
	r3 := multiply(12, 14, 22, -0)

	fmt.Println("1*4*8*6= ", r1)
	fmt.Println("645*897*385*(-66)= ", r2)
	fmt.Println("12*14*22*(-0)= ", r3)
}

func multiply(n ...int) int {
	result := 1
	for _, v := range n {
		if v == 0 {
			return 0
		} else {
			result *= v
		}
	}
	return result
}
