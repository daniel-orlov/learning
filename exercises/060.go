package main

//Create a user defined struct with
	//the identifier “person”
	//the fields:
		//first
		//last
		//age
//attach a method to type person with
	//the identifier “speak”
	//the method should have the person say their name and age
//create a value of type person
//call the method from the value of type person

import "fmt"

type person struct {
	first, last string
	age int
}

func (p person) speak() {
	fmt.Printf("My name is %v %v, I am %v y.o.", p.first, p.last, p.age)
}

func main()  {
	p0 := person{
		first: "Neo",
		last:  "Anderson",
		age:   33,
	}
	p0.speak()
}
