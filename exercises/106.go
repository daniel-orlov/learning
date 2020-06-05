package main

/*
Show the comma ok idiom
 */

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 777
	}()

	v, ok := <-c
		fmt.Println(v, ok)

	close(c)

	v, ok = <-c
		fmt.Println(v, ok)
}
