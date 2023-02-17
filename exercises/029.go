package main

//Continue using slices (slicing a slice, using slicing for iterating over a slice, appending to a slice)

import "fmt"

func main() {
	cf_fav_letters := []string{"H", "W", "P", "O"}
	//Slicing a slice
	first_two := cf_fav_letters[:2]
	two_in_the_middle := cf_fav_letters[1:3]
	last_two := cf_fav_letters[2:]

	fmt.Println(first_two)
	fmt.Println(two_in_the_middle)
	fmt.Println(last_two)
	//using slicing for iterating over a slice
	for i := 0; i <= 3; i++ {
		fmt.Println(i, cf_fav_letters[i])
	}
	//appending to a slice
	cf_fav_letters = append(cf_fav_letters, "!")
	fmt.Println(cf_fav_letters)
	full_version := []string{"Hard", "Work", "Pays", "Off", "!"}
	cf_fav_letters = append(cf_fav_letters, full_version...) //three dots to take all the elements
	//fmt.Println(cf_fav_letters)
	fmt.Println(cf_fav_letters[5:])
}
