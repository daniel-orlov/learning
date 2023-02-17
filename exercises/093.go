package main

import "fmt"

/*
Channels: creating a send-only and a receive-only channels
 */

func main()  {
	cs := make(chan <- int, 1) //send-only channel

	cs <- 128500
	/* THIS CODE FAILS
	fmt.Println(<- cs)
	fmt.Println(<- cs)
	*/

	fmt.Println("-----")
	fmt.Printf("%T\n", cs)

	cr := make(<- chan int, 1) //receive-only channel

	/* THIS CODE FAILS
	cs <- 128500

	fmt.Println(<- cr)
	*/


	fmt.Println("-----")
	fmt.Printf("%T\n", cr)

}