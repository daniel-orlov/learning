package main

import "fmt"

func main() {
	galeraBuilder := getShipBuilder("Galera")
	falconXBuilder := getShipBuilder("FalconX")

	spaniard := newDirector(galeraBuilder)
	typicalGalera := spaniard.BuildShip()
	fmt.Println(typicalGalera.Represent())

	genius := newDirector(falconXBuilder)
	falconXDD := genius.BuildShip()
	fmt.Println(falconXDD.Represent())
}
