package main

//Continue using slices (indices, len(), iterating through a slice)

import "fmt"

func main() {
	cf_fav_nums := []int{21, 15, 9}
	fmt.Printf("Slice includes %v elements\n", len(cf_fav_nums))
	fmt.Println(cf_fav_nums[2])
	//iterative printing of index-value pairs
	for i, v := range cf_fav_nums {
		fmt.Println(i, v)
	}
}
