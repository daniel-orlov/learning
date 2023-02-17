package main

/*
Fix the race condition you created in the previous exercise
by using a mutex
it makes sense to remove runtime.Gosched()
*/

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	incrementedValue := 0
	const goRouNum = 12
	var wg sync.WaitGroup
	wg.Add(goRouNum)
	var mut sync.Mutex

	for i := 0; i < goRouNum; i++ {
		go func() {
			mut.Lock()
			val := incrementedValue
			val++
			incrementedValue = val
			mut.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Printf("One goroutine should be left beyond this point: %v\n", runtime.NumGoroutine() == 1)
	fmt.Println(incrementedValue)
}
