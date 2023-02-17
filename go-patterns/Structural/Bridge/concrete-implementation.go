package main

import "fmt"

type hp struct {
}

func (p *hp) printFile() {
	fmt.Print("HP printer goes brr...\n\n")
}

type epson struct {
}

func (p *epson) printFile() {
	fmt.Print("Epson printer goes bzz...\n\n")
}
