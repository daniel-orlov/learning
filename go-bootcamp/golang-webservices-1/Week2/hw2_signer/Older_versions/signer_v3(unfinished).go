package main

import (
	"fmt"
	"sort"
	"strconv"
	// "sync"
	"bufio"
	"os"
	"strings"
	// "runtime"
	es "github.com/pkg/errors"
)

func md5Func(md5Done chan<- struct{}, in <-chan string, out chan<- string) {
	data := <-in
	out <- DataSignerMd5(data)
	// close(out)
	md5Done <- struct{}{}
}

func crc32Func(in <-chan string, out chan<- string) {
	data := <-in
	out <- DataSignerCrc32(data)
	// close(out)
}

func SingleHash(md5Done chan struct{}, in <-chan string, out chan<- string) {
	outMd5 := make(chan string)
	outCrc32 := make(chan string)
	outCrc32Md5 := make(chan string)

	go md5Func(md5Done, in, outMd5)
	go crc32Func(in, outCrc32)
	go crc32Func(outMd5, outCrc32Md5)

	out <- fmt.Sprintf("%v~%v", <-outCrc32, <-outCrc32Md5)
	close(out)
}

func MultiHash(in <-chan string, multiHashOut chan<- string) {
	data := <-in
	multiHashIn := make(chan string, 6)
	for i := 0; i < 6; i++ {
		thString := strconv.Itoa(i)
		multiHashIn <- thString + data
	}
	close(multiHashIn)
	for _ = range multiHashIn {
		go crc32Func(multiHashIn, multiHashOut)
	}
	// close(multiHashOut)
}

func CombineResults(in <-chan string) {
	buffer := make([]string, 0, 100)
	for i := range in {
		buffer = append(buffer, i)
	}
	sort.Slice(buffer, func(i, j int) bool { return buffer[i] < buffer[j] })
	var combinedResult string
	for _, i := range buffer {
		combinedResult += i + "_"
	}
	combinedResult = combinedResult[:(len(combinedResult) - 1)]
	fmt.Printf("CombineResults %v\n\n", combinedResult)
}

func ExecutePipeline(md5Done chan struct{}, in <-chan string, output chan<- string) {
	singleHashOut := make(chan string)
	go SingleHash(md5Done, in, singleHashOut)
	multiHashOut := make(chan string, 6)
	go MultiHash(singleHashOut, multiHashOut)
	var pipelineResult string
	for i := range multiHashOut {
		pipelineResult += i
	}
	output <- pipelineResult
	close(output)
}

func acceptInput() chan string {
	fmt.Println("Enter numeric values separated by spaces and press 'Enter'")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		err = es.Wrap(err, "failed to readfrom input")
		return err
	}
	input := make([]string, 0, 100)
	input = strings.Split(text, " ")
	size := len(input)

	out := make(chan string, 2*size)
	for _, i := range input {
		out <- i
		out <- i
	}
	close(out)
	return out
}

func main() {
	md5Done := make(chan struct{})
	pipelineInput := acceptInput()
	pipelineOutput := make(chan string)
	go ExecutePipeline(md5Done, pipelineInput, pipelineOutput)
	for _ = range pipelineInput {
		select {
		case <-md5Done:
			go ExecutePipeline(md5Done, pipelineInput, pipelineOutput)
			// runtime.Schedule()
		}
	}
	CombineResults(pipelineOutput)
}
