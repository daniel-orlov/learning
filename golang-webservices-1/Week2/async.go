package main

import (
	"fmt"
	"time"
)

func getComments() chan string {
	//use buffered chan - otherwise ther's risk of blocking -> leakage
	result := make(chan string, 1)
	go func(out chan<- string) {
		time.Sleep(2 * time.Second)
		fmt.Println("async operation is ready, return comments")
		out <- "32 comments"
	}(result)
	return result
}

func getPage() {
	resultChan := getComments()

	time.Sleep(1 * time.Second)
	fmt.Println("get related articles")

	// return

	commentsData := <-resultChan
	fmt.Println("main goroutine:", commentsData)
}

func main() {
	for i := 0; i < 3; i++ {
		getPage()
	}
}
