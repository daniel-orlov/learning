package main

import(
	"fmt"
	"strconv"
	"sort"
	"sync"
	// "runtime"
)

func SingleHash() string {
	var data string
	fmt.Scanln(&data)
	md5 := DataSignerMd5(data)  //cannot be multiplexed
	wg := &sync.WaitGroup{}

	crc32Chan := make(chan string)
	wg.Add(1)
	go func(data string, result chan string) {
		result <- DataSignerCrc32(data)
		wg.Done()
	}(data, crc32Chan)
	crc32 := <- crc32Chan

	crc32md5Chan := make(chan string)
	wg.Add(1)
	go func(data string, result chan string) {
		result <- DataSignerCrc32(data)
		wg.Done()
	}(md5, crc32md5Chan)
	crc32md5 := <- crc32md5Chan

	wg.Wait()

	singHsh := crc32 + "~" + crc32md5

	fmt.Println("md5:", md5)
	fmt.Println("crc32:", crc32)
	fmt.Println("crc32md5:", crc32md5)
	fmt.Println("singHsh:", singHsh)

	return singHsh
}

func MultiHash(th int, singHsh string) string {
	var data string
	data = strconv.Itoa(th) + singHsh
	crc32 := DataSignerCrc32(data)
	fmt.Println(th, " crc32:", crc32)
	return crc32
}

func main() {
	step1 := SingleHash()
	step2 := make([]string, 0, 6)
	multiHshChan := make(chan string, 6)
	wg := &sync.WaitGroup{}
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func(i int, step1 string, multiHshChan chan string){
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

	// fmt.Println("UNSORTED:", step2)
	sort.Slice(step2, func(i, j int) bool {return step2[i] < step2[j]})
	// fmt.Println("SORTED", step2)
	var result string
	for i := range step2 {
		result += step2[i]
	}
	fmt.Println("RESULT", result)

}
