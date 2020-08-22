package main

import "fmt"

func main() {
	//creating a slice
	var slice0 []int 					//len=0, cap=0, not init. (nil)
	fmt.Println(len(slice0), cap(slice0))

	slice1 := []int{}					//len=0, cap=0, init
	fmt.Println(len(slice1), cap(slice1))

	slice2 := []int{42}					//len=1, cap=1
	fmt.Println(len(slice2), cap(slice2))

	slice3 := make([]int, 0)			//len=0, cap=0
	fmt.Println(len(slice3), cap(slice3))

	slice4 := make([]int, 5)			//len=5, cap=5
	fmt.Println(len(slice4), cap(slice4))

	slice5 := make([]int, 5, 10)		//len=5, cap=10
	fmt.Println(len(slice5), cap(slice5))

	//accessing elem-s using index
	var buffer = make([]int, 5)
	fmt.Println(buffer[2])

	//adding elem-s to the end
	buffer = append(buffer, 1, 2, 3, 4)
	fmt.Println(buffer)
	fmt.Println(len(buffer), cap(buffer))

	//adding a slice to another slice
	newBuffer := make([]int, 3)
	buffer = append(buffer, newBuffer...)
	fmt.Println(buffer)
	fmt.Println(len(buffer), cap(buffer))

	//slicing a slice
	var sliceForSlicing = []int{0,1,2,3,4,5,6,7,8,9}
	fmt.Println(sliceForSlicing)
	sl1 := sliceForSlicing[1:8]
	sl2 := sliceForSlicing[:4]
	sl3 := sliceForSlicing[5:]
	fmt.Println(sl1, sl2, sl3)
}
