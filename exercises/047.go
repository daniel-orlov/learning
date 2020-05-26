package main

//Take the code from the previous exercise, then store the values of type person in a map with the key of first name. Access each value in the map. Print out the values, ranging over the slice.

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
		last: "The Master",
		favFlavours: []string{
			"raspberry",
			"honey",
		},
	}

	friends := map[string]person{
		pers1.first: pers1,
		pers2.first: pers2,
	}

	for _,v := range friends {
		for _, temp := range v.favFlavours {
			fmt.Printf("%v %v loves ice-cream with %v flavour\n", pers1.first, pers1.last, temp)
		}
	}
}