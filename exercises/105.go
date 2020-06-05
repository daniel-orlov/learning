package main

import (
	"fmt"
)

/*
Communicating using two channels
*/

func main()  {
	signalC := make(chan bool)
	inputC := record(signalC)

	play(inputC, signalC)

	fmt.Println("My work here is done...")
}

//this one does not need a goroutine as it only communicates with the stdOut
func play(c <-chan string, s <-chan bool) {
	for{
		select {
		case letter := <-c:
			fmt.Println(letter)
		case <-s: return
		}
	}
}

//this one needs a goroutine to make channel work
func record(s chan<- bool) <-chan string {
	c := make(chan string)
	abc := []string {
		"а", "б", "в", "г", "д", "е", "ё", "ж", "з", "и", "й", "к", "л", "м", "н", "о", "п", "р", "с", "т", "у", "ф", "х", "ц", "ч", "ш", "щ", "ъ", "ы", "ь", "э", "ю", "я",
	}
	go func() {
		for _, i := range abc {
			c <- i
		}
		s <- true
		close(c)
	}()
	return c
}