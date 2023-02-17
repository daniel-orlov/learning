package main

//Trying bit shifting on a decimal

import "fmt"

func main() {
	fmt.Print("DEC\t\tBIN\t\tHEX\n")
	some_variable := 42
	fmt.Printf("%d\t\t%b\t\t%#x\n", some_variable, some_variable, some_variable)
	shifted := some_variable << 1
	fmt.Printf("%d\t\t%b\t\t%#x\n", shifted, shifted, shifted)
}
