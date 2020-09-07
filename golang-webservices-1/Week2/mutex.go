package main

import (
	"fmt"
	"sync"
)

func main() {
	var counters = map[int]int{}
	mu := &sync.Mutex{}
	for i := 0; i < 5; i++ {
		go func(counters map[int]int, num int, mu *sync.Mutex) {
			for j:= 0; j < 5; j++ {
				//protecting writing
				mu.Lock()
				counters[num*10+j]++
				mu.Unlock()
			}
		}(counters, i, mu)
	}
	fmt.Scanln()
	//protecting reading
	mu.Lock()
	fmt.Println("counters result", counters)
	mu.Unlock()
}
