package main

//Using defer statement

import "fmt"

func main() {
	//W/o defer statement the order is going to change
	defer postponed()
	first()
}

func postponed() {
	fmt.Println("FIRST!")
}

func first() {
	fmt.Println("No, I am the FIRST!")
}