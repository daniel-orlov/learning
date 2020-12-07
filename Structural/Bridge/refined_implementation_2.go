package main

import "fmt"

type linuxLaptop struct {
	printerGUI printer
}

func (l *linuxLaptop) print() {
	fmt.Println("Printing a file using a desktop GUI:")
	l.printerGUI.printFile()
}

func (l *linuxLaptop) setPrinter(p printer) {
	l.printerGUI = p
}
