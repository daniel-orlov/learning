package main

//Creating custom functions in Go

import "fmt"

func main() {
	greetings("Dan")
	greet0 := strGreetings("Huginn")
	fmt.Println(greet0)
	greet1, status := greetUser("Bullwinkle", "Thorwaldson")
	fmt.Println(greet1, status)
}

//Simple function
func greetings(name string) {
	fmt.Println("Hello,", name)
}

//Function returning a string
func strGreetings(name string) string {
	return fmt.Sprint("Greetings,", name)
}

//Function with multiple returns
func greetUser(first string, last string) (string, bool) {
	var greet = fmt.Sprint("Well met, ", first, " ", last)
	isOnline := true
	return greet, isOnline
}