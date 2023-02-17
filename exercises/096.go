package main

import "fmt"

/*
Select statement to orchestrate the communications between channels
 */

func main()  {
	even := make(chan int)
	odd_ := make(chan int)
	quit := make(chan int)

	//send
	go send(even, odd_, quit)

	//receive
	receive(even, odd_, quit)
}

func receive(e, o, q <- chan int)  {
	for {
		select {
		case v := <- e:
			fmt.Println("from 'even' channel:", v)
		case v := <- o:
			fmt.Println("from 'odd_' channel:", v)
		case v := <- q:
			fmt.Println("from 'quit' channel:", v)
			return
		}
	}
}

func send(e,o,q chan<- int)  {
	for i := 0; i < 30; i++ {
		if i % 2 == 0{
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}