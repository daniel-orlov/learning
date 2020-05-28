package main

//Creating a custom type and declaring a variable 'x' of this type as done before
//Declaring a variable 'y' of type 'int' and assigning a value to it
//by conversion of variable 'x' into type 'int'

import "fmt"

type question int

var x question
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
