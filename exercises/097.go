package main

import (
	"fmt"
	"sync"
)

/*
Fanning in
 */

func main()  {
	even  := make(chan int)
	odd   := make(chan int)
	fanin := make(chan int)

	//send
	go send(even, odd)

	//receive
	go receive(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("Exiting...")
}

func receive(even, odd <-chan int, fanin chan<- int)  {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}

func send(even, odd chan<- int)  {
	for i := 0; i < 13; i++ {
		if i % 2 == 0{
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}