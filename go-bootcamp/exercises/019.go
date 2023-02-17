package main

//If statements and not '!' operator practice

import "fmt"

func main() {
	if true {
		fmt.Println("So true")
	}
	if false {
		fmt.Println("Completely false")
	}
	if !true {
		fmt.Println("So not true")
	}
	if !false {
		fmt.Println("Not false at all")
	}
}
