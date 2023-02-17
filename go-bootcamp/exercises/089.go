package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
Using goroutines, create an incrementer program:
		have a variable to hold the incrementer value;
		launch a bunch of goroutines;
			each goroutine should:
				read the incrementer value;
					store it in a new variable
				yield the processor with runtime.Gosched();
				increment the new variable;
				write the value in the new variable ->
				back to the incrementer variable;

use waitgroups to wait for all of your goroutines to finish
the above will create a race condition.
Prove that it is a race condition by using the -race flag
*/

func main() {
	incrementedValue := 0
	const goRouNum = 12
	var wg sync.WaitGroup
	wg.Add(goRouNum)

	for i := 0; i < goRouNum; i++ {
		go func() {
			val := incrementedValue
			runtime.Gosched()
			val++
			incrementedValue = val
			wg.Done()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Printf("One goroutine should be left beyond this point: %v\n", runtime.NumGoroutine() == 1)
	fmt.Println(incrementedValue)
}
