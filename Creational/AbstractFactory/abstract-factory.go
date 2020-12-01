package main

import "fmt"

type EquipmentFactory interface {
	makeKettlebell() iKettlebell
	makeBarbell() iBarbell
}

func getEquipmentFactory(brand string) (EquipmentFactory, error) {
	switch brand {
	case "Rogue":
		return &rogue{}, nil
	case "UralGirya":
		return &uralGirya{}, nil
	default:
		return nil, fmt.Errorf("unsupported factory brand: %v", brand)
	}
}
