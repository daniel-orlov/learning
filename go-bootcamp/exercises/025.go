package main

//Creating a not-so-useful weather app draft

import "fmt"

func main() {
	var weather string
	switch weather {
	case "rainy":
		fmt.Println("Bring an umbrella with you")
	case "sunny":
		fmt.Println("Apply sunscreen and wear sun glasses")
	case "snowy":
		fmt.Println("Wear a hat")
	case "meteor_rain":
		fmt.Print("Welcome to Mars!")
	default:
		fmt.Println("Stay at home anyway")
	}
}
