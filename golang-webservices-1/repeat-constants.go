package main

import "fmt"

const pi = 3.14159265

//init group of consts w/ iota
const (
	_ = iota
	KB uint64 = 1 << (10 * iota)
	MB
)


const (
	//untyped const
	year = 2020
	yearTyped int = 2020
)

func main() {
	fmt.Println(MB) //1048576

	var month int32 = 8
	allowedOperation := month + year
	fmt.Println(allowedOperation)
	//not allowed
	//unallowedOperation = month + yearTyped

}
