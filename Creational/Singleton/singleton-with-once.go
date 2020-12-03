package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func getInstanceWithOnce() *Singleton {
	if Single == nil {
		once.Do(
			func() {
				fmt.Println("Initiating single instance for the first time.")
				Single = &Singleton{}
			})
	} else {
		fmt.Println("Single instance already exists.")
	}
	return Single
}
