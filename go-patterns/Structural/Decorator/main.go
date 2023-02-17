package main

import "fmt"

func main() {

	snickersCakeS := &SnickersCakeBase{}

	//Add caramel layer
	snickersCakeM := &caramelLayer{
		ketoCake: snickersCakeS,
	}

	//Add tomato topping
	snickersCakeL := &chocolateDecorations{
		ketoCake: snickersCakeM,
	}

	fmt.Printf(
		"The price of Snickers keto-friendly cake with caramel layer and chocolate decorations is %d\n",
		snickersCakeL.getPrice(),
	)
}
