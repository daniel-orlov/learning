package main

//Creating an anonymous function

import "fmt"

func withName(name string) {
	fmt.Println("I am a func with a name! The name is", name)
}

func main() {
	withName("Nymous")

	func() { //params go in these parens
		fmt.Println("I am so secret. I am anonymous")
	}() //args in these parens

}
