package main

//Assigning function as a value to a variable
//Functions in Go are the first-class citizens
//Which means they can be assign to a variable, passed into function, returned by another function and so on

import "fmt"

func main() {
	function := func(num int) {
		fmt.Println("Welcome back,", num)
	}
	function(47)

	theLuckyNum := returning777()
	fmt.Printf("%T\n", theLuckyNum)
	anotherLayer := theLuckyNum()
	fmt.Println(anotherLayer)
	//Alternatively we can pass func straight into fmt.Println()
	fmt.Println(theLuckyNum())
	//or even less layers (and worse readability) like that:
	fmt.Println(returning777()())
}

func returning777() func() int {
	return func() int {
		return 777
	}
}