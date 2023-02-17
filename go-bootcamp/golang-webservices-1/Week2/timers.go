package main

import (
	"fmt"
	"time"
)

func longSQLQuery() chan bool {
	ch := make(chan bool, 1)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- true
	}()
	return ch
}

func main() {
	// 1 - timeout, 3 - normal execution
	timer := time.NewTimer(3 * time.Second)
	select {
	case <-timer.C:
		fmt.Println("timer.C timeout")
	case <-time.After(time.Minute):
		//GC can't collect it until time runs out
		fmt.Println("time.After timeout")
	case result := <-longSQLQuery():
		// frees resources
		if !timer.Stop() {
			<-timer.C
		}
		fmt.Println("operation result:", result)
	}
}
