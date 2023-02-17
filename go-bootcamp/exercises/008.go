package main

//Printing a number in decimal, binary and hexadecimal formats

import "fmt"

func main() {
	num := 1618
	fmt.Printf("dec = %v\nbin = %b\nhex = %#x", num, num, num)
}
