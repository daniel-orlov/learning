package main

import "fmt"

func main(){
	//copying a slice
	var patientZero = []int{101100, 111001, 100100} 	//has len = 3, cap = 3

	//THE WRONG WAY
	buf := make([]int, 0)
	fmt.Println(len(buf), cap(buf))
	copied := copy(buf, patientZero) //only copies the smallest slice
	fmt.Println(copied) //is actually empty

	//THE RIGHT WAY
	newBuf := make([]int, len(patientZero), cap(patientZero))
	copy(newBuf, patientZero)
	fmt.Printf("actual copy of patientZero: %#v\n", newBuf)

	//one can copy inside existing slide
	ints := make([]int, 4)
	copy(ints[1:3], []int{66, 99}) //ints = [0, 66, 99, 0]
	fmt.Println(ints)

}
