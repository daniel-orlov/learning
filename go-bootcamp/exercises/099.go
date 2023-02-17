package main

/*
Fanning in (inspired by Rob Pike's "boring Ann & Joe" example)
*/

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	outputC := FanIn(SlowCount(20), SlowFibonacci(20))
	for i := 0; i < 42; i++ {
		fmt.Println(<-outputC)
	}
	fmt.Println("My work is done here...")
}

func SlowCount(limit int) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i <= limit; i++ {
			c <- fmt.Sprint("Counting... ", i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func SlowFibonacci(limit int) <-chan string {
	c := make(chan string)
	go func() {
		x := 0
		y := 1
		z := 0
		for i := 0; i <= limit; i++ {
			z = x + y
			c <- fmt.Sprint("Fibonaccing ", z)
			x = y
			y = z
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func FanIn(i1, i2 <-chan string) <-chan string {
	output := make(chan string)
	go func() {
		for {
			output <- <-i1
		}
	}()
	go func() {
		for {
			output <- <-i2
		}
	}()
	return output
}
