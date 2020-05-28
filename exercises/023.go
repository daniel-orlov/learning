package main

//Printing out the remainder for each number between 10 and 100 when it is divided by 4.

import "fmt"

func main() {
	d := 4
	for i := 10; i < 101; i++ {
		fmt.Printf("%v is divided by %v. The remainder is %v\n", i, d, i%4)
	}
}
