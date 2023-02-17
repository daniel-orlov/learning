package main

import "fmt"

func getMeleeWeapon(weaponType string) (Melee, error) {
	switch weaponType {
	case "spear":
		return newSpear(), nil
	case "halberd":
		return newHalberd(), nil
	default:
		return nil, fmt.Errorf("unsupported weapon type: %v", weaponType)
	}
}
