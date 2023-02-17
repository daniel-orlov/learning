package main

import (
	"fmt"
)

type S1 struct{}

//Shadowed method
func (s1 S1) f() {
	fmt.Println("S1.f()")
}
func (s1 S1) g() {
	fmt.Println("S1.g()")
}

type S2 struct {
	S1
}

//Shadowing method
func (s2 S2) f() {
	fmt.Println("S2.f()")
}

type I interface {
	f()
}

func printType(i I) {
	fmt.Printf("%T\n", i)
	if s1, ok := i.(S1); ok {
		s1.f()
		s1.g()
	}
	if s2, ok := i.(S2); ok {
		s2.f()
		s2.g() //this one doesn't exist, so s1.g() is executed instead
	}
}

func main() {
	printType(S1{})
	printType(S2{})
}
