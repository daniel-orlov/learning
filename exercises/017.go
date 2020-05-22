package main

//Running 'for' loop, breaking and continue using conditional statement
//Printing only even numbers, as defined by conditional statement using modulus operator

import "fmt"

func main() {
	x := 0
	for {
		x++
		if x > 12 {
			break
		}
		if x % 2 != 0 {
			continue
		}
		fmt.Println(x)
	}
	fmt.Println("I am done with this loop... Yet again")
}
