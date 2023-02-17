package main

import "fmt"

/*
get this code working using a buffered channel
 */

/*BEFORE
func main()  {
	c := make(chan int)
	c <- 42
	fmt.Println(<-c)
}
*/

//AFTER

func main()  {
	c := make(chan int, 3)
	c <- 42
	c <- 43
	c <- 44
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}