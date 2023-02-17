package main

import "fmt"

func main() {
	fmt.Println("creating a slice with len = 5, cap = 5")
	var sl1 = make([]int, 5)
	fmt.Println("length = ", len(sl1))
	fmt.Println("capacity = ", cap(sl1))
	fmt.Println(sl1, "\n")

	fmt.Println("making a copy of it")
	newSlice := sl1[:]
	fmt.Println("length = ", len(newSlice))
	fmt.Println("capacity = ", cap(newSlice))
	fmt.Println(newSlice, "\n")

	fmt.Println("now let's replace the 2nd elem with 42")
	sl1[1] = 42
	fmt.Println(sl1)
	fmt.Println(newSlice, "\n")
	//both slices change,
	//because they are pointing to the same area in memory

	fmt.Println("then let's overflow the capacity of newSlice")
	newSlice = append(newSlice, 1969)
	fmt.Println("length = ", len(newSlice))
	fmt.Println("capacity = ", cap(newSlice))
	fmt.Println(newSlice, "\n")

	fmt.Println("finally let's modify sl1 and see if it still influences newSlice")
	sl1[3] = 77
	fmt.Println(sl1)
	fmt.Println(newSlice, "\n")
	//it doesn't anymore, as it is now in a different memory location
}
