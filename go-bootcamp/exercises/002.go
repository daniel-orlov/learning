package main

//Initiating variables in the package (global) scope:
//Declaring them with var
//Printing out 'zero values'

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
