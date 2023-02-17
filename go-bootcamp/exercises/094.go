package main

import "fmt"

/*
Using two specific channels to create a data flow.
Or having a chicken, at least.
*/

func main() {
	c := make(chan string)
	fmt.Println("Alright Chaps, LET'S DO THIS!")
	go sender(c)
	receiver(c)
	fmt.Println("JEEEENKIIIINSSS!!!!")
}

func sender(c chan<- string) {
	c <- "LEEEEEEEEEEERRRRROOOOOOOOOOY!!!!"
}

func receiver(c <-chan string) {
	fmt.Println(<-c)
}
