package main

import "fmt"

func deferDemo1() {
	// 1) Some usefull stuff
	// 2) Then defers in the reversed order

	defer fmt.Println("second - counting function running time")
	defer fmt.Println("first - closing a file")
	fmt.Println("Some usefull stuff being done here...")
}

func deferDemo2() {
	// """1) Second defer arguments are calculated when it's being called
	// prints 'getSomeVars execution'
	// 2) Some usefull stuff
	// 3) Actually prints the return of getSomeVars
	// prints 'getSomeVars return'
	// 4) Executes the first defer:	prints 'All work has been finished'

	defer fmt.Println("All work has been finished")
	defer fmt.Println(getSomeVars())
	fmt.Println("Some usefull stuff being done here...")
}

func deferDemo3() {
	// 1) Some usefull stuff
	// 2) Then defers in the reversed order
	// prints 'getSomeVars execution'
	// prints 'getSomeVars return'
	// prints 'All work has been finished'

	defer fmt.Println("All work has been finished")
	defer func() {
		fmt.Println(getSomeVars())
	}()
	fmt.Println("Some usefull stuff being done here...")
}

func panicDemo1() {
	//no panic handling
	fmt.Println("Some usefull stuff being done here...")
	panic("I sense the disturbance in the Force")
	return
}

func panicDemo2() {
	//defer still runs even in panic mode
	defer func() {
		fmt.Println("Still works in panic mode")
	}()
	fmt.Println("Some usefull stuff being done here...")
	panic("I sense the disturbance in the Force")
	return
}

func panicDemo3() {
	//handling panic with deferred recover
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic happened:", err)
		}
	}()
	fmt.Println("Some usefull stuff being done here...")
	panic("I sense the disturbance in the Force")
	return
}

func getSomeVars() string {
	fmt.Println("getSomeVars execution")
	return "getSomeVars return"
}

func main() {
	// deferDemo1()
	// deferDemo2()
	// deferDemo3()
	// panicDemo1()
	// panicDemo2()
	panicDemo3()
}
