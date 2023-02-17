package main

//Running 'for' loop, breaking using conditional statement

import "fmt"

func main() {
	x := 0
	for {
		if x > 5 {
			break
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("I am done with this loop...")
}
