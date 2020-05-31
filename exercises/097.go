package main

import (
	"fmt"
	"sync"
)

/*
Fanning in
 */

func main()  {
	even_ := make(chan int)
	odd__ := make(chan int)
	fanin := make(chan int)

	//send
	go send(even_, odd__)

	//receive
	go receive(even_, odd__, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("Exiting...")
}

func receive(even_, odd__ <-chan int, fanin chan<- int)  {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even_ {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd__ {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}

func send(even_, odd__ chan<- int)  {
	for i := 0; i < 13; i++ {
		if i % 2 == 0{
			even_ <- i
		} else {
			odd__ <- i
		}
	}
	close(even_)
	close(odd__)
}