package main

import "fmt"

type poleArm struct {
	name    string
	power   int
	agility int
}

func (p *poleArm) setName(name string) {
	p.name = name
}

func (p *poleArm) setPower(power int) {
	p.power = power
}

func (p *poleArm) setAgility(agility int) {
	p.agility = agility
}

func (p *poleArm) getName() string {
	return p.name
}

func (p *poleArm) getPower() int {
	return p.power
}

func (p *poleArm) getAgility() int {
	return p.agility
}

func (p *poleArm) represent() string {
	return fmt.Sprintf(
		"This is a %v (%T):\nPower:\t\t%v\nAgility:\t%v\n",
		p.name, p, p.power, p.agility,
	)
}
