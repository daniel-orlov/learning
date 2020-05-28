package main

//Using this slice:
//x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
//create the following new slices which are then printed (by SLICING):
//[42 43 44 45 46]
//[47 48 49 50 51]
//[44 45 46 47 48]
//[43 44 45 46 47]

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	s0 := x[:5]
	s1 := x[5:]
	s2 := x[2:7]
	s3 := x[1:6]

	fmt.Printf("%v\n%v\n%v\n%v\n", s0, s1, s2, s3)
}
