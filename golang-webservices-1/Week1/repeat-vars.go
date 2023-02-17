package main

import (
	"fmt"
)

func main() {
	//explicit type
	var num0 int = 10
	fmt.Println(num0)

	//implicit type
	var num1 = 20
	fmt.Println(num1)

	//short init operator
	num2 := 30
	fmt.Println(num2)

	//+= operator
	num2 += 10
	fmt.Println(num2)

	//post-incrementing with ++
	num1++
	fmt.Println(num1)
}
