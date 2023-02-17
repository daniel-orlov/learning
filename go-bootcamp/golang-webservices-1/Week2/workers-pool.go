package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

const goroutinesNum = 3

func startWorker(workerNum int, in <-chan string) {
	for input := range in {
		fmt.Printf(formatWork(workerNum, input))
		runtime.Gosched() // try commenting out this line
	}
	printFinishWork(workerNum)
}

func main() {
	runtime.GOMAXPROCS(0)               // try 0 (all) and 1
	workerInput := make(chan string, 2) // try increasing chan size
	for i := 0; i < goroutinesNum; i++ {
		go startWorker(i, workerInput)
	}

	months := []string{"January", "February", "March",
		"April", "May", "June",
		"July", "August", "September",
		"October", "November", "December",
	}

	for _, monthName := range months {
		workerInput <- monthName
	}
	close(workerInput) // try commenting out this line

	time.Sleep(time.Millisecond)
}

func formatWork(in int, input string) string {
	return fmt.Sprintln(strings.Repeat("  ", in), "█",
		strings.Repeat("  ", goroutinesNum-in),
		"go", in,
		"recieved", input)
}

func printFinishWork(in int) {
	fmt.Println(strings.Repeat("  ", in), "█",
		strings.Repeat("  ", goroutinesNum-in),
		"===", in,
		"finished")
}
