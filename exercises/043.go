package main

//Creating a value of type person with underlying type struct, composing together values of different types

import "fmt"

type person struct {
	first, last string
	age         int
}

func main() {
	user1 := person{
		first: "Brent",
		last:  "Fikowski",
		age:   29,
	}
	user2 := person{
		first: "Pat",
		last:  "Vellner",
		age:   29,
	}
	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Printf("%v %v is %v years old.", user1.first, user1.last, user1.age)
}
