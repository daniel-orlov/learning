package main

//A “callback” is when we pass a func into a func as an argument. For this exercise, perform a callback.

import "fmt"

func main() {
	x := 47
	result := chooseDivider(dividesBy2, x)
	fmt.Println(result)
	result1 := chooseDivider(dividesBy3, x)
	fmt.Println(result1)
	result2 := chooseDivider(dividesBy4, x)
	fmt.Println(result2)
	result3 := chooseDivider(dividesBy47, x)
	fmt.Println(result3)
}

func chooseDivider(f func(int) bool, num int) bool {
	return f(num)
}

func dividesBy2(num int) bool{
	if num % 2 == 0 {
		return true
	}else{
		return false
	}
}

func dividesBy3(num int) bool {
	if num % 3 == 0 {
		return true
	}
	return false
}

func dividesBy4(num int) bool{
	if num % 4 == 0 {
		return true
	}else{
		return false
	}
}

func dividesBy47(num int) bool{
	if num % 47 == 0 {
		return true
	}else{
		return false
	}
}
