package main

import "fmt"

type windows struct {
}

func (w *windows) runWindowsOS() {
	fmt.Println("Running Windows 7 like there is no tomorrow.")
}
