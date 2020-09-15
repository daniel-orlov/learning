package main

import (
	"fmt"
	"strconv"
	"sort"
	"sync"
	"strings"
	"bufio"
	"os"
	// "runtime"
	es "github.com/pkg/errors"
)

const (
	MULTIHASHPARTS = 6
)

type Hash struct {
	Input 				string
	SingleHashStruct
	MultiHashStruct
}

type SingleHashStruct struct {
	Md5, Crc32, Crc32Md5, SingleHash string
}

type MultiHashStruct struct {
	MulCrc32 		[]string
	MulCrc32Str 	string
}

func md5Func(md5Done chan struct{}, in string, out chan string) {
	out <- DataSignerMd5(in)
	close(out)
	md5Done <- struct{}{}
}

func crc32Func(in string, out chan string) {
	out <- DataSignerCrc32(in)
	// close(out)
}

func SingleHash(md5Done chan struct{}, hash Hash) (<-chan Hash) {
	out := make(chan Hash, 10)
	outMd5 := make(chan string, 10)
	outCrc32 := make(chan string, 10)
	outCrc32Md5 := make(chan string, 10)

	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func(md5Done chan struct{}, in string, out chan string){
		md5Func(md5Done, hash.Input, outMd5)
		wg.Done()
	}(md5Done, hash.Input, outMd5)

	go func(in string, out chan string){
		crc32Func(hash.Input, outCrc32)
		wg.Done()
	}(hash.Input, outCrc32)

	go func(in string, out chan string){
		crc32Func(hash.Md5, outCrc32Md5)
		wg.Done()
	}(hash.Md5, outCrc32Md5)


	hash.Md5 = <-outMd5
	hash.Crc32 = <-outCrc32
	hash.Crc32Md5 = <-outCrc32Md5
	hash.SingleHash = fmt.Sprintf("%v~%v", hash.Crc32, hash.Crc32Md5)

	out <- hash
	wg.Wait()
	close(out)
	return out
}

func MultiHash(in <-chan Hash) (<-chan Hash) {
	out := make(chan Hash, 10)
	multiHashOut := make(chan string, MULTIHASHPARTS)
	results := make([]string, MULTIHASHPARTS)
	data := ""

	hash := <-in
	// wg := &sync.WaitGroup{}
	// wg.Add(MULTIHASHPARTS)
	for i := 0; i < MULTIHASHPARTS; i++ {
		data = strconv.Itoa(i) + hash.SingleHash
		// multiHashOut = make(chan string)
		// go func(in string, out chan string){
		crc32Func(data, multiHashOut)
			// wg.Done()
		// }(data, multiHashOut)
		//MUTEX HERE?
		results[i] = <- multiHashOut
	}
	for _, val := range results {
		hash.MulCrc32Str += val
	}


	hash.MulCrc32 = results
	out <- hash
	close(out)
	return out
}


func ExecutePipeline(md5Done chan struct{}, in string, out chan<- Hash)  {
	hash := Hash{Input: in,}
	multiHashOut := MultiHash(SingleHash(md5Done, hash))
	for i := range multiHashOut{
		out<- i
	}
	close(out)
}

func CombineResults(wg *sync.WaitGroup, in <-chan Hash){
	buffer := make([]string, 0, MaxInputDataLen)
	for i := range in {
		buffer = append(buffer, i.MulCrc32Str)
	}
	sort.Slice(buffer, func(i, j int) bool {return buffer[i] < buffer[j]})
	var combinedResult string
	for _, i := range buffer {
		combinedResult += i + "_"
	}
	// combinedResult = combinedResult[:(len(combinedResult)-1)]
	fmt.Printf("CombineResults %v\n\n", combinedResult)
	wg.Done()
}

func acceptInput() (chan string, error) {
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

func main() {
	in, err := acceptInput()
	if err != nil {
		err = es.Wrap(err, "Input failed")
		fmt.Println(err)
	}
	out := make(chan Hash, MaxInputDataLen)
	md5Done := make(chan struct{}, MaxInputDataLen)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go CombineResults(wg, out)

	for val := range in{
		wg.Add(1)
		go func(md5Done chan struct{}, val string, out chan Hash){
			ExecutePipeline(md5Done, val, out)
			wg.Done()
		}(md5Done, val, out)
		<-md5Done
	}
	wg.Wait()

}
