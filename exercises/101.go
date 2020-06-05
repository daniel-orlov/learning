package main

/*
get this code working using func literal, aka, anonymous self-executing func
*/

import "fmt"

/*BEFORE
func main()  {
	c := make(chan int)
	c <- 42
	fmt.Println(<-c)
}
 */

//AFTER

func main()  {
	c := make(chan int)
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}