package main

/*
Creating a new value of type error, panicking and printing the error out
 */

import (
	"errors"
	"fmt"
)

var IncorrectAnswer = errors.New("incorrect answer: we insist that you should agree with us")

func main()  {
	myPosition := "I challenge the status quo"
	reaction, err := socialPressure(myPosition)
	defer fmt.Println(reaction)
	if err != nil {
		panic(err)
	}

}

func socialPressure(a string) (string, error) {
	if a != "you all are right" {
		return "How dare you?!", IncorrectAnswer
	}
	return "Very well.", nil
}