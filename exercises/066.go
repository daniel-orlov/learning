package main

//Closure is when we have “enclosed” the scope of a variable in some code block. For this hands-on exercise, create a func which “encloses” the scope of a variable:

import "fmt"

func main()  {
	x := 47
	y := 12
	fmt.Println("This is main.func body comparison:", x > y)
	{
		y := 99
		fmt.Println("This is enclosed comparison:", x > y)
	}
}
