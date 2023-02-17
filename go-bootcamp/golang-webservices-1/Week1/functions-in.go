package main

import "fmt"

// simple declaration of a func
func singleInt(integer int) int {
	return integer
}

// func w/ multiple param-s
func multiIntSum(a, b int, c int) int {
	return a + b + c
}

// func w/ named return
func namedReturnDemo(x int) (output int) {
	output = x * 2
	return
}

// func w/ multiple returns
func multipleReturns(in int) (int, error) {
	if in == 0 {
		return in, fmt.Errorf("had zero as input")
	}
	return 2 / in, nil
}

// func w/ multiple named returns
func multipleNamedReturns(in int) (quatered int, err error) {
	if in == 0 {
		quatered = 0
		err = fmt.Errorf("had zero as input")
	}
	quatered = in / 4
	err = nil
	return
}

// func with variadic param-s
func sumAll(in ...int) (result int) {
	fmt.Printf("Input: %#v \n", in)
	for _, val := range in {
		result += val
	}
	return
}

func main() {
	fmt.Println(singleInt(12))
	fmt.Println(multiIntSum(12, 85, 0))
	fmt.Println(namedReturnDemo(12))
	fmt.Println(multipleReturns(12))
	fmt.Println(multipleNamedReturns(12))
	fmt.Println(sumAll(12, 21, 52, 5653, 63, 323, 512, 65))
}
