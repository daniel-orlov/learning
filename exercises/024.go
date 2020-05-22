package main

//A simple demo of mix of if/else if/else statements and conditional operators. With a flavour of global craze.

import "fmt"

func main () {
	var have_symptoms bool
	var days_passed int //since you are back from your trip

	have_symptoms = false
	days_passed = 84

	if have_symptoms {
		fmt.Println("Self-quarantine and notify your recent contacts, if any")
	} else if !have_symptoms && days_passed < 15 {
		fmt.Printf("Sorry, no meetings for you for at least %v days", 15 - days_passed)
	} else {
		fmt.Println("Stop reading Twitter")
	}
}