package main

//Create a func which returns a func
//assign the returned func to a variable
//call the returned func

import "fmt"

func main() {
	f := selector(true)
	fmt.Println(f())
}

//never ever call your func select, that cost me like 15 minutes to figure out what's wrong with my perfectly viable code

func selector(mode bool) func() string {
	if mode {
		return func() string {
			return "This Is Funcception"
		}
	}
	return selector(true)
}
