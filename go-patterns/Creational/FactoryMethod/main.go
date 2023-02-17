package main

import (
	"fmt"
	"os"
)

const (
	HALBERD = "halberd"
	SPEAR   = "spear"
	KHOPESH = "khopesh"
)

func main() {

	halberd, err := getMeleeWeapon(HALBERD)
	if err != nil {
		err = fmt.Errorf("unable to get a melee weapon: %w", err)
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(halberd.represent())
	}

	spear, err := getMeleeWeapon(SPEAR)
	if err != nil {
		err = fmt.Errorf("unable to get a melee weapon: %w", err)
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(spear.represent())
	}

	khopesh, err := getMeleeWeapon(KHOPESH)
	if err != nil {
		err = fmt.Errorf("unable to get a melee weapon: %w", err)
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(khopesh.represent())
	}

}
