package main

import(
	"fmt"
	"strconv"
	"sort"
	"sync"
	"strings"
	"bufio"
	"os"
)

func SingleHash(data string, singleHashChan chan<- string) {
	md5 := DataSignerMd5(data)  //cannot be multiplexed
	//TODO: pass signal to other pipelines that md5 is done
	wg := &sync.WaitGroup{}
	crc32Chan := make(chan string) //Computing crc32 below
	wg.Add(1)
	go func(data string, result chan string) {
		result <- DataSignerCrc32(data)
		wg.Done()
	}(data, crc32Chan)

	crc32md5Chan := make(chan string) //Computing crc32(md5) below
	wg.Add(1)
	go func(md5 string, result chan string) {
		result <- DataSignerCrc32(md5)
		wg.Done()
	}(md5, crc32md5Chan)

	crc32 := <- crc32Chan
	crc32md5 := <- crc32md5Chan
	singHsh := crc32 + "~" + crc32md5
	wg.Wait()
	fmt.Printf("%v SingleHash data %v\n", data, data) 	//Printing out all the values in order
	fmt.Printf("%v SingleHash md5(data) %v\n", data, md5)
	fmt.Printf("%v SingleHash crc32(md5(data)) %v\n", data, crc32md5)
	fmt.Printf("%v SingleHash crc32(data) %v\n", data, crc32)
	fmt.Printf("%v SingleHash result %v\n", data, singHsh)

	singleHashChan <- singHsh
}

func MultiHash(th int, singHsh string) []string {
	var data string
	thString := strconv.Itoa(th)
	data = thString + singHsh
	crc32 := make([]string, 0, 2)
	crc32 = append(crc32, thString)
	crc32 = append(crc32, DataSignerCrc32(data))
	return crc32
}


func CombineResults(size int, resultChan <-chan string){
	buffer := make([]string, 0, size)
	for i := range resultChan {
		buffer = append(buffer, i)
	}
	sort.Slice(buffer, func(i, j int) bool {return buffer[i] < buffer[j]})
	var combinedResult string
	for _, i := range buffer {
		combinedResult += i + "_"
	}
	combinedResult = combinedResult[:(len(combinedResult)-1)]
	fmt.Printf("CombineResults %v\n\n", combinedResult)
}



func ExecutePipeline(data string, pipelineChan chan string) {
	wg := &sync.WaitGroup{}

	//STEP 1: pass data into SingleHash
	singleHashChan := make(chan string)
	wg.Add(1)
	go func(data string, singleHashChan chan string) {
		SingleHash(data, singleHashChan)
		wg.Done()
	}(data, singleHashChan)
	step1 := <- singleHashChan
	close(singleHashChan)

	//STEP 2: multiplex MultiHash x 6
	step2 := make([][]string, 0, 6)
	multiHshChan := make(chan []string, 6)
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func(i int, step1 string, multiHshChan chan []string){
			result := MultiHash(i, step1)
			multiHshChan <- result
			wg.Done()
		}(i, step1, multiHshChan)
	}
	wg.Wait()
	close(multiHshChan)
	for i := range multiHshChan {
		step2 = append(step2, i)
	}

	sort.Slice(step2, func(i, j int) bool {return step2[i][0] < step2[j][0]})
	var multiHashResult string
	for i := range step2 { //Printing out all the values in order + creating result string
		fmt.Printf("%v MultiHash: crc32(th+step1)) %v %v\n", step1, step2[i][0], step2[i][1])
		multiHashResult += step2[i][1]
	}
	fmt.Printf("%v MultiHash result: %v \n", step1, multiHashResult)
	//pass result to CombinedResult
	pipelineChan <- multiHashResult
	// close(multiHshChan)
}

func main() {
	//Accepting input from console
	fmt.Println("Enter numeric values separated by spaces and press 'Enter'")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	checkPanic(err)
	input := make([]string, 0, 100)
	input = strings.Split(text, " ")
	size := len(input)
	//Executing pipelines
	pipelineChan := make(chan string, size)
	for _, data := range input {
		ExecutePipeline(data, pipelineChan)
	}
	close(pipelineChan)
	//printing combined result
	CombineResults(size, pipelineChan)
}

func checkPanic(err error) {
	if err != nil {
		panic(err)
	}
}
