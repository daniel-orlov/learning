package main

import (
	"fmt"
	"os"
)

const (
	POOD = 16.0
)

func main() {
	//Creating a concrete factory
	rogueFactory, err := getEquipmentFactory("Rogue")
	if err != nil {
		err = fmt.Errorf("unable to get equipment factory: %w", err)
		fmt.Fprintln(os.Stderr, err)
		return
	}

	//Using the factory to create concrete objects
	kbRogue := rogueFactory.makeKettlebell()
	fmt.Println(kbRogue.represent())
	bbRogue := rogueFactory.makeBarbell()
	fmt.Println(bbRogue.represent())

	//Creating a concrete factory
	uralGiryaFactory, err := getEquipmentFactory("UralGirya")
	if err != nil {
		err = fmt.Errorf("unable to get equipment factory: %w", err)
		fmt.Fprintln(os.Stderr, err)
		return
	}
	//Using the factory to create concrete objects
	kbUralGirya := uralGiryaFactory.makeKettlebell()
	fmt.Println(kbUralGirya.represent())
	bbUralGirya := uralGiryaFactory.makeBarbell()
	fmt.Println(bbUralGirya.represent())

	//Setters in action
	kbUralGirya.setColor("green")
	kbUralGirya.setWeight(POOD * 1.5)
	fmt.Println(kbUralGirya.represent())
	bbRogue.setWeight(10)
	bbRogue.setPurpose("Hafthor Bjornsson to use it as a selfie-stick")
	fmt.Println(bbRogue.represent())
}
