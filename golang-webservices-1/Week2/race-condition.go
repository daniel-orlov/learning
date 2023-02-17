package main

import "fmt"

func main() {
	var counters = map[int]int{}
	for i := 0; i < 5; i++ {
		go func(counters map[int]int, num int) {
			for j:= 0; j < 5; j++ {
				counters[num*10+j]++
			}
		}(counters, i)
	}
	fmt.Scanln()
	fmt.Println("counters result", counters)
}
