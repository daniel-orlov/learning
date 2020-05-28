package main

//Assign a func to a variable, then call that func

import "fmt"

func main() {
	v1 := 42
	v2 := 47
	result := func(v1 int, v2 int) {
		if v1 > v2 {
			fmt.Println(true)
		} else if v1 < v2 {
			fmt.Println(false)
		} else {
			fmt.Println(nil)
		}
	}
	result(v1, v2) //false
	f := result
	v3, v4 := 77, 69
	f(v3, v4) //true
	g := f
	v5 := v4
	g(v4, v5) //nil
}
