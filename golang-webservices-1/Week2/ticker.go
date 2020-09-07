package main

import(
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	i := 0
	for tickTime := range ticker.C {
		i++
		fmt.Println("step", i, "time", tickTime)
		if i >= 5 {
			//must initiate stop
			ticker.Stop()
			break
		}
	}
	fmt.Println("total", i)
	// return

	// time.Tick is an alias for Ticker, works infinitely
	// cannot be collected by garbage collector
	c := time.Tick(time.Second)
	i = 0
	for tickTime := range c {
		i++
		fmt.Println("step", i, "time", tickTime)
		if i >= 5 {
			//cannot be stopped, so we just break the loop
			break
		}
	}
}
