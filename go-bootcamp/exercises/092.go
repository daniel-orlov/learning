package main

import "fmt"

/*
Channels - channels block and two ways to overcome it
*/

func main() {
	/* THIS IS NOT GONNA RUN - CAUSING A DEADLOCK
	fatal error: all goroutines are asleep - deadlock!

	c := make(chan int)
	c <- 46
	fmt.Println(<- c)
	*/

	//This code runs - first solution, launching a goroutine
	c := make(chan int)
	go func() {
		c <- 47
	}()
	fmt.Println("CHANNEL C")
	fmt.Println(<-c)

	//This code runs - second solution, buffer channel
	c1 := make(chan int, 1) //this is an expandable buffer channel (currently holding only one value)

	c1 <- 48

	fmt.Println("CHANNEL C1")
	fmt.Println(<-c1)

	/* THIS IS NOT GONNA RUN - CAUSING A DEADLOCK
		fatal error: all goroutines are asleep - deadlock!

	c2 := make(chan int, 1) //this is an expandable buffer channel (currently holding ONLY ONE value)
		c2 <- 49
		c2 <- 50 //this crashes the buffer channel
		fmt.Println("CHANNEL C2")
	fmt.Println(<- c2)
	*/

}
