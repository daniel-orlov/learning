package main

//Printing out all ASCII Characters corresponding to dec from 65 to 122

import "fmt"

func main() {
	for x:= 65; x <= 122; x++ {
		fmt.Printf("%v\t%#U\n", x, x)
	}
}