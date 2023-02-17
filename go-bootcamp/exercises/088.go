package main

import "fmt"

/*
This exercise will reinforce our understanding of method sets:

create a type person struct
attach a method speak to type person using a pointer receiver
		*person

create a type human interface
		to implicitly implement the interface,
		a human must have the speak method

create func “saySomething”
		have it take in a human as a parameter
		have it call the speak method

show the following in your code
		you CAN pass a value of type *person into saySomething
		you CANNOT pass a value of type person into saySomething
*/

type person struct {
	first, last string
}

func (p *person) speak(sp string) {
	fmt.Printf("%v %v says: %v", p.first, p.last, sp)
}

type human interface {
	speak(sp string)
}

func saySomething(h human, sp string) {
	h.speak(sp)
}

func main() {
	//you CAN pass a value of type *person into saySomething
	AL := person{
		first: "Abraham",
		last:  "Lincoln",
	}
	saySomething(&AL, "Whose boots do you think I should black?")

	//you CANNOT pass a value of type person into saySomething
	//saySomething(AL, "Whose boots do you think I should black?")
	//This code above doesn't run
}
