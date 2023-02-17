package main

import (
	"bufio"
	es "github.com/pkg/errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	MULTIHASHPARTS 	= 6
	BUFFER 			= 10	//This one is arbitrary
)

type Hash struct {
	Input 								string
	Md5, Crc32, Crc32Md5, SingleHash 	string
	MulCrc32    						[][]string
	MulCrc32Str 						string
}

func main() {
	defer elapsed("MAIN")() //HERE AND BELOW: Optional measuring time of execution
	in, err := acceptInput()
	if err != nil {
		err = es.Wrap(err, "Input failed")
		fmt.Println(err)
	}
	expectedLoad := len(in)
	out := make(chan Hash, expectedLoad)
	md5Done := make(chan struct{}, BUFFER)
	//Listening for incoming results of the pipelines
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go CombineResults(expectedLoad, wg, out)
	//Launching pipelines
	for val := range in {
		wg.Add(1)
		go func(md5Done chan struct{}, val string, out chan Hash) {
			ExecutePipeline(md5Done, val, out)
			wg.Done()
		}(md5Done, val, out)
		<-md5Done //Signal for new goroutine to start
	}
	wg.Wait()
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func ExecutePipeline(md5Done chan struct{}, in string, out chan<- Hash) {
	// defer elapsed("ExecutePipeline")()
	hash := Hash{Input: in}
	multiHashOut := MultiHash(SingleHash(md5Done, hash))
	for i := range multiHashOut {
		out <- i
	}
}

func SingleHash(md5Done chan struct{}, hash Hash) <-chan Hash {
	/*Computes SingleHash in parallel and passes the result to MultiHash
	*/
	// defer elapsed("SingleHash")()
	out := make(chan Hash, BUFFER)
	outMd5 := make(chan string, BUFFER)
	outCrc32 := make(chan string, BUFFER)
	outCrc32Md5 := make(chan string, BUFFER)

	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func(in string, out chan string, md5Done chan struct{}) {
		md5Func(hash.Input, outMd5, md5Done)
		wg.Done()
	}(hash.Input, outMd5, md5Done)

	go func(in string, out chan string) {
		crc32SingleFunc(hash.Input, outCrc32)
		wg.Done()
	}(hash.Input, outCrc32)

	hash.Md5 = <-outMd5
	go func(in string, out chan string) {
		crc32SingleFunc(hash.Md5, outCrc32Md5)
		wg.Done()
	}(hash.Md5, outCrc32Md5)

	hash.Crc32 = <-outCrc32
	hash.Crc32Md5 = <-outCrc32Md5
	hash.SingleHash = fmt.Sprintf("%v~%v", hash.Crc32, hash.Crc32Md5)
	out <- hash
	wg.Wait()
	close(out)
	return out
}

func MultiHash(in <-chan Hash) <-chan Hash {
	/*Computes MultiHash in parallel and passes the result back to ExecutePipeline
	*/
	// defer elapsed("MultiHash")()
	out := make(chan Hash, BUFFER)
	multiHashOut := make(chan []string, MULTIHASHPARTS)
	results := make([][]string, MULTIHASHPARTS)
	combinedString := ""

	hash := <-in
	wg := &sync.WaitGroup{}
	wg.Add(MULTIHASHPARTS)
	for i := 0; i < MULTIHASHPARTS; i++ {
		combinedString = strconv.Itoa(i) + hash.SingleHash
		go func(i int, combinedString string, out chan []string) {
			crc32MultiFunc(i, combinedString, multiHashOut)
			wg.Done()
		}(i, combinedString, multiHashOut)
	}

	for i := 0; i < MULTIHASHPARTS; i++ {
		temp := <-multiHashOut
		index, err := strconv.Atoi(temp[0])
		if err != nil {
			err = es.Wrap(err, "strconv.Atoi failed")
			fmt.Println(err)
		}
		results[index] = temp
	}

	for _, val := range results {
		hash.MulCrc32Str += val[1]
	}

	hash.MulCrc32 = results
	out <- hash
	close(out)
	return out
}

func md5Func(in string, out chan string, md5Done chan struct{}) {
	/*Computes md5 and sends signal upon finishing to avoid overheat
	*/
	// defer elapsed("md5Func")()
	out <- DataSignerMd5(in)
	close(out)
	md5Done <- struct{}{}
}

func crc32SingleFunc(in string, out chan string) {
	/*Computes crc32 as a part of SingleHash function
	*/
	// defer elapsed("crc32Func")()
	out <- DataSignerCrc32(in)
}

func crc32MultiFunc(i int, in string, out chan []string) {
	/*Computes crc32 as a part of MultiHash function
	*/
	// defer elapsed("crc32Func")()
	crc32 := make([]string, 2)
	crc32[0] = strconv.Itoa(i)
	crc32[1] = DataSignerCrc32(in)
	out <- crc32
}


func CombineResults(expected int, wg *sync.WaitGroup, in <-chan Hash) {
	/*Listens for the pipeline results, stores them, sorts and prints out
	*/
	// defer elapsed("CombineResults")()
	buffer := make([]string, 0, MaxInputDataLen)
	var result Hash
	for i := 0; i < expected; i++ {
		result = <-in
		buffer = append(buffer, result.MulCrc32Str)
	}
	sort.Slice(buffer, func(i, j int) bool { return buffer[i] < buffer[j] })
	var combinedResult string
	for index, val := range buffer {
		if index != len(buffer)-1 {
			combinedResult += val + "_"
			continue
		}
		combinedResult += val
	}
	fmt.Printf("CombineResults: %v\n\n", combinedResult)
	wg.Done()
}

func acceptInput() (chan string, error) {
	/*Accepts input from std.In? parses it and puts on a channel
	that will be fed to the pipelines
	*/
	defer elapsed("acceptInput")()
	fmt.Println("Enter numeric values separated by spaces and press 'Enter'")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		err = es.Wrap(err, "failed to read from reader")
		return nil, err
	}
	input := make([]string, 0, MaxInputDataLen)
	input = strings.Split(text, " ")
	size := len(input)

	out := make(chan string, size)
	for _, i := range input {
		out <- i
	}
	close(out)
	return out, err
}
