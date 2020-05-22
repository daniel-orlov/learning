package main

//Using comparison operators to initialize boolean type variables

import "fmt"

func main(){
	g := (1 == 0)
	h := (24 <= 42)
	i := (77 >= 77)
	j := (12 != 21)
	k := (69 > 96)
	l := (1.618 < 3.14159265)

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", g, h, i, j, k, l)
}
