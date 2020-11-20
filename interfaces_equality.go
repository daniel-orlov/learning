package main

import (
	"fmt"
	"reflect"
)

type S struct {
	a, b, c string
}

func main() {
	fmt.Println("WITH DIFFERENT STRUCTS")
	x := interface{}(&S{"a", "b", "c"})
	fmt.Printf("Value 'x': %#v\n", x)   //Value 'x': &main.S{a:"a", b:"b", c:"c"}
	fmt.Printf("Value '*x': %#v\n", &x) //Value '*x': (*interface {})(0xc000010200)
	y := interface{}(&S{"a", "b", "c"})
	fmt.Printf("Value 'y': %#v\n", y)   //Value 'y': &main.S{a:"a", b:"b", c:"c"}
	fmt.Printf("Value '*y': %#v\n", &y) //Value '*y': (*interface {})(0xc000010210)

	fmt.Println("ORDINARY EQUALITY CHECK")
	fmt.Println(x == y) //false

	fmt.Println("DEEP EQUALITY CHECK")
	fmt.Println(reflect.DeepEqual(x, y))

	fmt.Println("WITH THE SAME STRUCT")
	s3 := S{"a", "b", "c"}

	x = interface{}(&s3)
	fmt.Printf("Value 'x': %#v\n", x)
	fmt.Printf("Value '*x': %#v\n", &x)
	y = interface{}(&s3)
	fmt.Printf("Value 'y': %#v\n", y)
	fmt.Printf("Value '*y': %#v\n", &y)
	fmt.Println(x == y) //false
}
