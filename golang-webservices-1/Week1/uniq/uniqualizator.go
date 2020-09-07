package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

func unique(input io.Reader, output io.Writer) error {
	//reading from reader
	in := bufio.NewScanner(input)
	var prev string
	for in.Scan() {
		text := in.Text()
		if text == prev {
			continue
		} else if text < prev {
			return fmt.Errorf("file is not sorted")
		}
		prev = text
		fmt.Fprintln(output, text)
	}
	return nil
}


func main() {
	err := unique(os.Stdin, os.Stdout)
	if err != nil {
		panic(err.Error())
	}
}
