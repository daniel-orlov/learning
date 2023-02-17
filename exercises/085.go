package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
Given the code from the previous ex.,
Fixing a data race
*/

func main()  {
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	counter := 0

	const gsNum = 100
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(gsNum)

	for i := 0; i < gsNum; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("counted till: ", counter)
}