package main

import "fmt"

// a simple func
func doNothing() {
	fmt.Println("Today I swear I am not gonna do anything...")
}

func main() {
	//anonymous function
	func(input string) {
		fmt.Println("anon func output: ", input)
	}("ha-ha, classics")

	//assigning function to a variable
	double := func(input int) {
		fmt.Println("double input =", input*2)
	}

	double(2)

	//creating a type of func based on its signature
	type intToStrFuncType func(int) string
	type simpleIntFuncType func(int)
	type simpleStringFuncType func(string)

	//func taking other func as param (CALLBACK)
	executor := func(callback simpleIntFuncType) {
		callback(100)
	}

	executor(double)

	//function returning another function
	prefixer := func(prefix string) simpleStringFuncType {
		return func(input string) {
			fmt.Printf("[%s] %s\n", prefix, input)
		}
	}

	successLogger := prefixer("SUCCESS")
	successLogger("everything worked fine")

}
