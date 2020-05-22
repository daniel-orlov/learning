package main

//Creating constants: a typed one and an untyped one

import "fmt"

const (
	typed_one int = 42
	untyped_one = 1.618
)

func main()	{
	fmt.Printf("constant of type '%T' with value %v\n", typed_one, typed_one)
	fmt.Printf("constant of type '%T' with value %v", untyped_one, untyped_one)
}