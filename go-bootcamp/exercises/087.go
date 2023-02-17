package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
in addition to the main goroutine,
launch two more goroutines
each additional goroutine should print something out
use waitgroups to make sure each goroutine finishes
before your program exists
*/

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go GoFunc("First")
	go GoFunc("Second")

	wg.Wait()
	fmt.Println("By this time all but one goroutines should be finished.")
	fmt.Println("Number of goroutines: ", runtime.NumGoroutine())
}

func GoFunc(prefix string) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v counting: %v\n", prefix, i)
	}
	wg.Done() //Stuff is done, we can stop waiting
}
