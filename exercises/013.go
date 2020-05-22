package main

//Using iota to calculate years I need to master Go
//JK, it will probably take less/more time

import "fmt"

const (
	current_year = 2020 + iota
	one_year_from_now = current_year + iota
	two_years_from_now = current_year + iota
	three_years_from_now = current_year + iota
	four_years_from_now = current_year + iota
)

func main() {
	fmt.Println(current_year)
	fmt.Println(one_year_from_now)
	fmt.Println(two_years_from_now)
	fmt.Println(three_years_from_now)
	fmt.Println(four_years_from_now)
}