package main

/*
Fanning out to handle multiple tasks simultaneously and then fanning in to present the results
*/

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go PopulateWithRandInts(c1)

	go FanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("I am so done with that... Exiting")
}

func FanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func(n int) {
			c2 <- TimeConsumingWork(n)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func TimeConsumingWork(n int) int {
	fmt.Println("I must consult the Elder Gods...")
	a := rand.Intn(10)
	time.Sleep(time.Second * time.Duration(a))
	return n * a
}

func PopulateWithRandInts(c chan int) {
	for i := 0; i <= 20; i++ {
		c <- rand.Intn(1e3)
	}
	close(c)
}
