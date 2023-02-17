package main

import (
	"fmt"
	"sync"
)

var mtx = &sync.Mutex{}

type Singleton struct{}

var Single *Singleton

func getInstanceWithMtx() *Singleton {
	if Single == nil {
		mtx.Lock()
		defer mtx.Unlock()
		fmt.Println("Initiating single instance for the first time.")
		Single = &Singleton{}
	} else {
		fmt.Println("Single instance already exists.")
	}
	return Single
}
