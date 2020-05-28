package main

//create a type SQUARE
//create a type CIRCLE
//attach a method to each that calculates AREA and returns it
//circle area= π r 2
//square area = L * W
//create a type SHAPE that defines an interface as anything that has the AREA method
//create a func INFO which takes type shape and then prints the area
//create a value of type square
//create a value of type circle
//use func info to print the area of square
//use func info to print the area of circle

import (
	"fmt"
	"math"
)

type square struct {
	sideLength float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

type shape interface {
	area() float64
}

func main() {
	red := square{sideLength: 1147.0}
	angel := circle{radius: 6.66}
	a0 := info(red)
	a1 := info(angel)
	fmt.Printf("The area of the %T is %v\n", red, a0)
	fmt.Printf("The area of the %T is %v\n", angel, a1)
}

func info(s shape) float64 {
	return s.area()
}
