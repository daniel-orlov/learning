package main

//Learning how to use maps and their features + "comma ok" idiom

import "fmt"

func main() {
	pers_bests := map[string]int{
		"Deadlift":     180,
		"Strict Press": 60,
		"Strict MU":    0,
	}
	fmt.Println(pers_bests)
	fmt.Println(pers_bests["Deadlift"])
	//Adding a new k-v pair - easily
	pers_bests["TGU"] = 40
	fmt.Println(pers_bests)
	//Deleting a k-v pair - easily (even if it doesn't exist!)
	delete(pers_bests, "Strict MU")
	fmt.Println(pers_bests)
	delete(pers_bests, "Strict BMU") //As I said - even if it doesn't exist
	fmt.Println(pers_bests)

	fmt.Println(pers_bests["DUs"]) //doesn't throw a key error! Prints '0' instead
	//Using "comma ok" idiom to check if the key is actually in the map
	fmt.Println("Comma OK idiom below")
	v, ok := pers_bests["DUs"]
	fmt.Println(v)
	fmt.Println(ok)
	//Common use of the principle above
	if v, ok := pers_bests["DUs"]; ok {
		fmt.Println("Value for DUs PR is", v) //this doesn't get printed
	}
	if v, ok := pers_bests["Strict Press"]; ok {
		fmt.Println("The value for Strict Press PR was", v, "kilos") //this gets printed; also that's before quarantine, so probably outdated
	}
}
