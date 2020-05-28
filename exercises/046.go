package main

//Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
//first name
//last name
//favorite ice cream flavors
//Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.

import (
	"fmt"
)

type person struct {
	first, last string
	favFlavours []string
}

func main() {
	pers1 := person{
		first: "Huginn",
		last:  "The Raven",
		favFlavours: []string{
			"blueberry",
			"chocolate",
		},
	}
	pers2 := person{
		first: "Pooh",
		last:  "The Master",
		favFlavours: []string{
			"raspberry",
			"honey",
		},
	}
	for _, temp := range pers1.favFlavours {
		fmt.Printf("%v %v loves ice-cream with %v flavour\n", pers1.first, pers1.last, temp)
	}
	for _, temp := range pers2.favFlavours {
		fmt.Printf("%v %v loves ice-cream with %v flavour\n", pers2.first, pers2.last, temp)
	}
}
