package main

import "fmt"

type androidPhone struct {
	printerApp printer
}

func (a *androidPhone) print() {
	fmt.Println("Printing a file using a native mobile app:")
	a.printerApp.printFile()
}

func (a *androidPhone) setPrinter(p printer) {
	a.printerApp = p
}
