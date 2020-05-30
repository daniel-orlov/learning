package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/*
Same problem being solved as in the previous ex.,
This time using the pkg sync/atomic
*/

func main()  {
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	var counter int64 //important for operations using atomic pkg

	const gsNum = 100
	var wg sync.WaitGroup
	wg.Add(gsNum)

	for i := 0; i < gsNum; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Count: ", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("counted till: ", counter)
}