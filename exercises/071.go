package main

/*
Using label to mark the chunk of code as Loop
We can also control the code flow using goto Loop
*/

import "fmt"

func main() {
	x := 12

InfNestedLoop:
	for {
		for {
			x++
			fmt.Println(x)
			switch x {
			case 21:
				fmt.Println("Target value reached, breaking")
				break InfNestedLoop //breaking the whole lable InfNestedLoop
			}
		}
	}

}
