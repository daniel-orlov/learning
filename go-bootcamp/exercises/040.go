package main

//Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:
//
//"James", "Bond", "Shaken, not stirred"
//"Miss", "Moneypenny", "Helloooooo, James."
//
//Range over the records, then range over the data in each record.

import "fmt"

func main() {
	jb := []string{
		"James", "Bond", "Shaken, not stirred",
	}
	fmt.Println(jb)
	mp := []string{
		"Miss", "Moneypenny", "Helloooooo, James.",
	}
	fmt.Println(mp)

	quotes := [][]string{
		jb, mp,
	}
	fmt.Println(quotes)

	for ko, _ := range quotes {
		for ki, vi := range quotes[ko] {
			fmt.Printf("Coll.%d, item%d - %v\n", ko, ki, vi)
		}
	}
}
