package main

type spear struct {
	poleArm
}

func newSpear() Melee {
	return &spear{poleArm{
		name:    "spear",
		power:   76,
		agility: 64,
	}}
}

type halberd struct {
	poleArm
}

func newHalberd() Melee {
	return &halberd{poleArm{
		name:    "halberd",
		power:   92,
		agility: 43,
	}}
}
