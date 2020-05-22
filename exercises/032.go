package main

//Creating a multi-dimensional slice

import "fmt"

func main() {
	header := []string{"NAME", "SURNAME", "COUNTRY", "NICKNAME"}
	fmt.Println(header)
	bf := []string{"Brent", "Fikowski", "Canada", "the Professor"}
	fmt.Println(bf)
	pv := []string{"Patrick", "Vellner", "Canada", "Uncle Pat"}
	fmt.Println(pv)
	table := [][]string{header, bf, pv}
	fmt.Println(table)
}
