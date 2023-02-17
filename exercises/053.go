package main

//Interface practice

import "fmt"

type person struct {
	first, last string
	age         int
}

type robot struct {
	serialNumber string
	needs        []string
}

type secretAgent struct {
	person
	inLove bool
}

type sentient interface {
	introduce()
}

func (s secretAgent) introduce() {
	fmt.Printf("My name is %v. %v %v.\n", s.last, s.first, s.last)
}

func (r robot) introduce() {
	fmt.Printf("I am %v. I need your %v and your %v\n", r.serialNumber, r.needs[0], r.needs[1])
}

func talk(s sentient) {
	fmt.Println("I can talk!")
}

func main() {
	sa007 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		inLove: true,
	}
	fmt.Println(sa007)
	sa007.introduce()
	talk(sa007)

	t800 := robot{
		serialNumber: "T-800",
		needs: []string{
			"clothes",
			"motorbike",
		},
	}
	t800.introduce()
	talk(t800)
}
