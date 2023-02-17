package main

import "fmt"

//Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string which stores their favorite things. Store three records in your map. Print out all of the values, along with their index position in the slice.

func main() {
	favourites := map[string][]string{
		`bond_james`: []string{
			`Shaken, not stirred`, `Martinis`, `Women`,
		},
		`moneypenny_miss`: []string{
			`James Bond`, `Literature`, `Computer Science`,
		},
		`no_dr`: []string{
			`Being evil`, `Ice cream`, `Sunsets`,
		},
	}
	for k, _ := range favourites {
		fmt.Printf("Record for %v:\n", k)
		for i, iv := range favourites[k] {
			fmt.Printf("\t\t\t%d - %v\n", i, iv)
		}
	}
}
