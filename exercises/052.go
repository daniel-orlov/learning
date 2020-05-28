package main

//Adding methods to

import "fmt"

type person struct {
	first, last string
	age         int
}

type secretAgent struct {
	person
	inLove bool
}

func (s secretAgent) introduce() {
	fmt.Printf("My name is %v. %v %v.", s.last, s.first, s.last)
}

func main() {
	sa007 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		inLove: true,
	}
	fmt.Println(sa007)
	sa007.introduce()
}
