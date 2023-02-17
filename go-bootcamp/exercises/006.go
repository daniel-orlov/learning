package main

//Getting the OS and Architecture type printed by Go

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

}
