package main

//Printing every rune code point of the uppercase alphabet two times, separated in pairs

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		for j := 0; j <= 1; j++ {
			fmt.Printf("%#U\n", i)
		}
		fmt.Println("")
	}
}