package main

import "fmt"

//Using the code from the previous example, add a record to your map and delete a record. Now print the map out using the “range” loop

func main() {
	favourites := map[string][]string{
		`bond_james`:  []string{
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
	//adding a record
	fmt.Printf("\nAdding a record...\n\n")
	favourites["Q"] = []string{
		"Backend development", "Manuals", "Exploding stationary",
	}
	//printing out
	for k, _ := range favourites {
		fmt.Printf("Record for %v:\n", k)
		for i, iv := range favourites[k] {
			fmt.Printf("\t\t\t%d - %v\n", i, iv)
		}
	}
	//deleting a record
	fmt.Printf("\nDeleting a record...\n\n")
	delete(favourites, "no_dr" )
	//printing out
	for k, _ := range favourites {
		fmt.Printf("Record for %v:\n", k)
		for i, iv := range favourites[k] {
			fmt.Printf("\t\t\t%d - %v\n", i, iv)
		}
	}
}
