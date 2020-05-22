package main

//Initiating variables as done previously
//"Sprinting" them to a string and assigning it to the variable 's'

import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main() {
	s := fmt.Sprintln(x, y, z)
	fmt.Print(s)
}