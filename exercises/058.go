package main

//create a func with the identifier foo that returns an int
//create a func with the identifier bar that returns an int and a string
//call both funcs
//print out their results

import "fmt"

func main() {
	result0 := foo()
	result1, result2 := bar()
	fmt.Println(result0, result1, result2)
}

func bar() (int, string) {
	return (foo() + 10005) , "this is a BAR"
}

func foo(...int) int {
	return 2340
}