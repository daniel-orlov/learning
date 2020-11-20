package main

import (
	"fmt"
	"sync"
)

const MAGIC = 10

func main() {
	m1 := make(map[int]int)
	m2 := make(map[int]int)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	wg.Add(2 * MAGIC)

	//Loop without feeding i into go func
	for i := 0; i < MAGIC; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			m1[i] = i
			mu.Unlock()
		}()
	}

	//Loop with feeding i into go func
	for i := 0; i < MAGIC; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m2[i] = i
			mu.Unlock()
		}(i)
	}
	wg.Wait()

	//See the difference
	fmt.Printf("Without feeding:\t%#v\n", m1)
	fmt.Printf("With feeding:\t\t%#v\n", m2)

}
