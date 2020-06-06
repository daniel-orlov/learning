package main

import (
	"fmt"
	"time"
)

/*
Create a struct “customErr” which implements the builtin.error interface.
Create a func “foo” that has a value of type error as a parameter.
Create a value of type “customErr” and pass it into “foo”.
 */

type customErr struct{
	num int
	dsc string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Struct customErr #%v caused the following error: %v", ce.num, ce.dsc)
}

func main()  {
	ce := customErr{
		num: 1984,
		dsc: "Unable to find your location",
	}
	foo(ce)
}

func foo(err error)  {
	fmt.Println("Doing something time consuming...")
	time.Sleep(4*time.Duration(time.Second))
}