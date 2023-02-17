package main

/*
Using comma ok idiom to check when a channel is closed
 */

import "fmt"

func main()  {
	even := make(chan int)
	odd  := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)

	receive(even, odd, quit)

}

func send(even, odd chan<- int, quit chan<- bool)  {
	for i := 0; i < 13; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

func receive(even, odd <-chan int, quit <-chan bool)  {
	for  {
		select {
		case v := <-even:
			fmt.Println("A value coming from even chan: ", v)
		case v := <-odd:
			fmt.Println("A value coming from odd chan: ", v)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("From comma ok: ", i, ok)
				return
			} else {
				fmt.Println("From comma ok: ", i)
			}
		}
	}
}