package main

import "fmt"

/*
Exploring another way to code a func that computes factorials
 */

//Initiating a global var to control the func
var globalFactorialController int = 7

func main()  {
//Computing and printing all factorials from 0 upto controller
	num := 0
	for num <= globalFactorialController {
		fmt.Printf("Factorial !%d ", num)
		result := 1
		counter := num

	Loop:
		if counter == 0 {
			goto Result
		}
		result = counter * (result)
		counter--
		goto Loop

	Result:
		fmt.Printf("= %v\n", result)
		num++
	}
}