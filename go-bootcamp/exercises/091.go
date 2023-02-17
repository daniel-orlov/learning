package main

/*
Fix the race condition you created in the previous exercise
by using package atomic
*/

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var incrementedValue int64
	const goRouNum = 12
	var wg sync.WaitGroup
	wg.Add(goRouNum)

	for i := 0; i < goRouNum; i++ {
		go func() {
			atomic.AddInt64(&incrementedValue, 1)
			fmt.Println(atomic.LoadInt64(&incrementedValue))
			wg.Done()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Printf("One goroutine should be left beyond this point: %v\n", runtime.NumGoroutine() == 1)
	fmt.Println(incrementedValue)
}
