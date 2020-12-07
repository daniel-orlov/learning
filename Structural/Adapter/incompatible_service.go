package main

import "fmt"

type linux struct{}

func (l *linux) runLinuxOS() {
	fmt.Println("Running your favourite Linux distributive.")
}
