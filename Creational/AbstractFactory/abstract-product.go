package main

type iKettlebell interface {
	setColor(color string)
	setWeight(weight float64)
	getColor() string
	getWeight() float64
	represent() string
}

type kettlebell struct {
	color  string
	weight float64
}

type iBarbell interface {
	setPurpose(purpose string)
	setWeight(weight float64)
	getPurpose() string
	getWeight() float64
	represent() string
}

type barbell struct {
	purpose string
	weight  float64
}
