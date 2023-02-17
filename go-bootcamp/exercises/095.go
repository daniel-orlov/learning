package main

import "fmt"

/*
Ranging over a channel and closing it after it's done
*/

func main() {
	c := make(chan int)
	fmt.Println("For science!")
	//SEND
	go func() {
		for i := 20; i >= 0; i-- {
			c <- i
		}
		close(c) //if we don't close - deadlock
	}()
	//RECEIVE
	for val := range c {
		fmt.Println(val)
	}
	fmt.Println("LAUNCH")
}
