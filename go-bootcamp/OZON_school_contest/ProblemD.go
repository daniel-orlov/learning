package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {
	var a, b string
	fmt.Println("Enter two numbers you want to sum-up:")
	fmt.Fscanf(os.Stdin, "%v %v", &a, &b)
	var bigA, _ = new(big.Int).SetString(a, 10)
	var bigB, _ = new(big.Int).SetString(b, 10)
	var product = new(big.Int)
	product.Add(bigA, bigB)
	fmt.Printf("The result is: %v", product)
}
