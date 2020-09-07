package main

import "fmt"

func main() {
	a := 42
	b := &a	//now b holds a pointer to a
	*b = 69 //now a = 69
	c := &a //we created a new pointer to a

	d := new(int) //pointer to a var with type int (zero default value)
	*d = 12
	*c = *d // c = 12 -> a = 12 ('a' has changed its value)
	*d = 21 // c and a are intact

	c = d //now c holds the same pointer as d
	*c = 777 //now c = 777 -> d = 777, a = 12
	fmt.Printf("a = %v, b = %v, c = %v, d = %v\n", a, *b, *c, *d)

}
