package main

//Creating a custom type
//Declaring a variable of the custom type

import "fmt"

type question int
var x question

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}