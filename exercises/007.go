package main

//Learning two ways how to calculate bits:
//with and without iota
//Meanwhile also learning about bit shifting

import "fmt"

//with iota
const (
	_ = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

//without iota
//const (
//	kb = 1024
//	mb = 1024 * kb
//	gb = 1024 * mb
//)

func main() {
	fmt.Printf("%v\t\t\t%b\n", kb, kb)
	fmt.Printf("%v\t\t\t%b\n", mb, mb)
	fmt.Printf("%v\t\t%b\n", gb, gb)
}