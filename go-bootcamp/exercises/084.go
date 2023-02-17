package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
Creating a data race - not a good thing - for learning purposes
*/

func main() {
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	counter := 0

	const gsNum = 100
	var wg sync.WaitGroup
	wg.Add(gsNum)

	for i := 0; i < gsNum; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("counted till: ", counter)
}
