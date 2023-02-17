package main

import "fmt"

func main() {
	s := "123" //the original string (OrStr), IMMUTABLE

	ps := &s //pointer to the OrStr

	b := []byte(*ps) //[]byte (MUTABLE) made with a pointer to an OrStr
	//NB: this is a new one, NOT the underlying one for the OrStr

	pb := &b //pointer to the new []byte

	s += "4"   //concatenation to the OrStr
	*ps += "5" //concatenation to the OrStr (this time addressing it with a pointer)
	b[1] = '0' //mutation of []byte: changing the 1-th index char to be equal '0'

	fmt.Println("Printed using a pointer to the OrStr: \t\t", *ps)
	fmt.Println("Printed using a pointer to the new []byte: \t", string(*pb))
}
