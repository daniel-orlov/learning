package main

import "fmt"

type vmAdapter struct {
	linuxImage *linux
}

func (a *vmAdapter) runWindowsOS() {
	fmt.Println("Preparing  a Linux VM to run on Windows")
	a.linuxImage.runLinuxOS()
}
