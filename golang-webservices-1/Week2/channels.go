package main

import "fmt"

func main() {
	// ch1 := make(chan int) //unbuffered channel

	ch1 := make(chan int, 1)

	go func(in chan int) {
		val := <- in
		fmt.Println("GO: received from channel 1", val)
		fmt.Println("GO: after receipt")
	}(ch1)

	ch1 <- 42
	ch1 <- 47

	fmt.Println("MAIN: after sending to channel 1")
	fmt.Scanln()
}
