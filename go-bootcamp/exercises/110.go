package main

/*
Including more information in the error using fmt pkg
*/

import (
	"fmt"
)

func main() {
	var joke string = "I don't trust stairs. They are always up to something."
	result, err := func(j string) (string, error) {
		if j != "" {
			return "", fmt.Errorf("OBS! Joke is already non-empty: %v", j)
		}
		return "Do you know what's brown and sticky? A stick", nil
	}(joke)
	if err != nil {
		panic(err)
	}
	defer fmt.Println(result)
}
