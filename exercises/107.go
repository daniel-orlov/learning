package main

/*
write a program that:
	puts 100 numbers to a channel
	pull the numbers off the channel and print them
 */

import "fmt"

func main()  {
	s := make(chan bool)

	c := putNums(s)

	pullNums(c, s)

	fmt.Println("Piece of cake is a lie")
	close(s)
}

func pullNums(c <-chan int, s <-chan bool) {
	for{
		select {
		case num := <-c:
			fmt.Println("Incoming: ", num)
		case <-s:
			return
		}
	}
}

func putNums(s chan<- bool) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++{
			c <- i
		}
		s <- true
		close(c)
	}()
	return c
}


