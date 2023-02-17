package main

//Trying different types of 'for'-statements:
// - with single condition;
// - with for clause;
// - with range clause;
// -

import "fmt"

func main() {
	//for statement with single condition
	x := 1
	for x <= 3 {
		fmt.Println(x)
		x++
	}
	//for statement with for clause
	//for init; condition; post{ S() }
	for y := 3; y > 0; y-- {
		fmt.Println(y)
	}
	//for statement with for clause

}
