package main

type Melee interface {
	setName(name string)
	setPower(power int)
	setAgility(agility int)
	getName() string
	getPower() int
	getAgility() int
	represent() string
}
