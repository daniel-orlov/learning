package main

import "fmt"

func main() {
	cancelChan := make(chan bool)
	dataChan := make(chan int)

	go func(cancelChan chan bool, dataChan chan int){
		val := 0
		for {
			select {
			case <-cancelChan:
				return
			case dataChan <- val:
				val++
			}
		}
	}(cancelChan, dataChan)

	for currentValue := range dataChan {
		fmt.Println("received: ", currentValue)
		if currentValue > 5 {
			fmt.Println("sending cancel")
			cancelChan <- true
			break
		}
	}

}
