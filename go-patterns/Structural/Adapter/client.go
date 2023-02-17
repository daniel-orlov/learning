package main

import "fmt"

type Client struct {
}

func (c *Client) runOS(hw hardware) {
	fmt.Println("Running the OS on the hardware:")
	hw.runWindowsOS()
}

type hardware interface {
	runWindowsOS()
}
