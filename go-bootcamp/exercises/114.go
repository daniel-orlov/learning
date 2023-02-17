package main

/*
Given this code,
use the sqrt.Error struct as a value of type error
*/

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		e := fmt.Errorf("more coffee needed - value was %v", f)
		return 0, sqrtError{"64°35'00.3\"N", "17°06'49.1\"W", e}
	}
	return 42, nil
}
