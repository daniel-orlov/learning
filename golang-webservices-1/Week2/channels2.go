package main

import "fmt"

func main() {
	in := make(chan int)

	go func(out chan<- int) {
		for i := 0; i <= 4; i++ {
			fmt.Println("before sending", i)
			out <- i
			fmt.Println("after sending", i)
		}
		close(out)
		fmt.Println("Generator has finished")
	}(in)

	for i := range in {
		fmt.Println("\treceived:", i)
	}

	// fmt.Scanln()
}
