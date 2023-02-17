package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
Go routines - INTRO
 */



var wg sync.WaitGroup

func main()  {
	fmt.Println("OS\t\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t\t", runtime.GOARCH)
	fmt.Println("Number of CPUs\t\t", runtime.NumCPU())
	fmt.Println("Number of goroutines\t", runtime.NumGoroutine())

	wg.Add(1)
	go GoFunc("Still") //Launching a goroutine. KA-CHOOUU!
	SimpleFunc()
	fmt.Println("Number of CPUs\t\t", runtime.NumCPU())
	fmt.Println("Number of goroutines\t", runtime.NumGoroutine())
	wg.Wait() //Prevents program from immediately exiting
}

func GoFunc(prefix string)  {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v counting: %v\n", prefix, i)
	}
	wg.Done() //Stuff is done, we can stop waiting
}

func SimpleFunc()  {
	for i := 1; i <= 10; i++ {
		fmt.Println("Counting: ", i)
	}
}