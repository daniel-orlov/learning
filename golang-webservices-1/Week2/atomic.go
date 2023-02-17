package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var totalOperations int32 = 0

func increment() {
	atomic.AddInt32(&totalOperations, 1)
}

func main() {
	for i := 0; i < 1000; i++{
		go increment()
	}
	time.Sleep(2 * time.Millisecond)
	//expecting totalOperations = 1000, but in fact...
	fmt.Println("totalOperations = ", totalOperations)
}
