package main

//Creating for loops, trying nesting them
//The plan was to create a vague representation of a chess board,
//but I do not yet know how to create a string-containing iterable in Go...
//TODO: learn how to crate lists/list-like structures and iterate over them

import "fmt"

func main() {
	for num := 1; num <= 8; num++ {
		fmt.Println()
		for letter := 1; letter <= 8; letter++ {
			fmt.Printf("%d:%d ", num, letter)
		}
	}
}
