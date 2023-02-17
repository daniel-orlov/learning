package main

//Create and use an anonymous struct.

import "fmt"

func main() {
	wod := struct {
		name, sort      string
		content         map[string][]int
		equipmentNeeded []string
	}{
		name: "Fran",
		sort: "for time",
		content: map[string][]int{
			"pull-ups": []int{
				21, 15, 9,
			},
			"thrusters": []int{
				21, 15, 9,
			},
		},
		equipmentNeeded: []string{
			"bar", "barbell", "plates",
		},
	}
	fmt.Println(wod)
}
