package main

//Embedding one type inside the other

import "fmt"

type person struct{
	first, last string
	age int
}

type crossfitter struct {
	person
	background string
}

func main() {
	cf1 := crossfitter {
		person: person {
			first: "Brent",
			last: "Fikowski",
			age: 29,
		},
		background: "volleyball",
	}
	cf2 := crossfitter{
		person: person {
			first: "Pat",
			last: "Vellner",
			age: 29,
		},
		background: "gymnastics",
	}


	fmt.Println(cf1)
	fmt.Println(cf2)
	fmt.Printf("%v %v is %v years old, has background in %v\n", cf1.first, cf1.last, cf1.age, cf1.background)
	fmt.Printf("%v %v is %v years old, has background in %v\n", cf2.first, cf2.last, cf2.age, cf2.background)
}