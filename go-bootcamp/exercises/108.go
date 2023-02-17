package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
write a program that:
		launches 10 goroutines
			each goroutine adds 10 numbers to a channel
		pull the numbers off the channel and print them
*/

const goRoutinesNum = 10

func main() {
	c := make(chan int)

	go quickRandPopulate(c)

	for nums := range c {
		fmt.Println(nums)
	}

	fmt.Println("....")
	fmt.Println("DONE")
}

func quickRandPopulate(c chan<- int) {
	var wg sync.WaitGroup

	for i := 0; i < goRoutinesNum; i++ {
		wg.Add(1)
		go func(m int) {
			for j := 0; j < 10; j++ {
				c <- rand.Intn(1e5) * m
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	close(c)
}
