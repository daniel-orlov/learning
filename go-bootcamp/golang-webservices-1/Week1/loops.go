package main

import "fmt"

func main() {
	//no condition whatsoever, potentially infinite (like while(true))
	for {
		fmt.Println("we are stuck here forever, unless...")
		break //You did it. You crazy son of a ...
	}

	//no condition, unless value changes
	// this is also potentially infinite (like while(someBool))
	someBool := true
	for someBool {
		fmt.Println("Poom-poom-poom. Another one stucks in loop...")
		//unless...
		someBool = false
	}

	//condition with init block
	for i := 0; i < 4; i++ {
		fmt.Println("Count: ", i)
	}

	//iterating over a slice
	someSlice := []int{12, 85, 0, 0}
	index := 0

	//first option - while style loop
	fmt.Println("WHILE STYLE LOOP")
	for index < len(someSlice) {
		fmt.Println("index:", index, "value:", someSlice[index])
		index++
	}

	//second option - C-style loop
	fmt.Println("C STYLE LOOP")
	for i := 0; i < len(someSlice); i++ {
		fmt.Println("index:", i, "value:", someSlice[i])
	}

	//third option - iteration by index and/or value
	fmt.Println("ITERATING OVER SLICE BY INDEX")
	for idx := range someSlice {
		fmt.Println("range slice by index", idx)
	}
	fmt.Println("ITERATING OVER SLICE BY INDEX AND VALUE")
	for idx, val := range someSlice {
		fmt.Println("range slice by index", idx, "and value", val)
	}

	//iterating over map
	count := map[int]string{1: "uno", 2: "dos", 3: "tres"}
	fmt.Println("ITERATING OVER MAP")

	//by key
	for key := range count {
		fmt.Println("range map by key", key)
	}

	//by value
	for _, val := range count {
		fmt.Println("range map by key", val)
	}

	//by key and value
	for key, val := range count {
		fmt.Println("range map by key", key, val)
	}

	//iterating over a string ([]byte)
	str := "slice of byte"
	fmt.Println("ITERATING OVER A STRING ([]BYTE)")
	for pos, char := range str {
		fmt.Printf("%#U at pos %d\n", char, pos)
	}
}
