package main

import "fmt"

/*
Using newly learned doc string
and using func init() to initialize a map
before func main() execution
*/

var globalMap map[string]string

func init() {
	globalMap = make(map[string]string)
	globalMap["benchmark"] = "Fran"
}

func main() {
	fmt.Println(globalMap) //it is already non-empty
}
