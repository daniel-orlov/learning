package main

import "fmt"

type rogueKettlebell struct {
	kettlebell
}

func (kb *rogueKettlebell) setColor(color string) {
	kb.color = color
}

func (kb *rogueKettlebell) setWeight(weight float64) {
	kb.weight = weight
}

func (kb *rogueKettlebell) getColor() string {
	return kb.color
}

func (kb *rogueKettlebell) getWeight() float64 {
	return kb.weight
}

func (kb *rogueKettlebell) represent() string {
	return fmt.Sprintf(
		"This is a %v %v-kilo kettlebell (%T)\n",
		kb.color, kb.weight, kb,
	)
}

type rogueBarbell struct {
	barbell
}

func (bb *rogueBarbell) setPurpose(purpose string) {
	bb.purpose = purpose
}

func (bb *rogueBarbell) getPurpose() string {
	return bb.purpose
}

func (bb *rogueBarbell) setWeight(weight float64) {
	bb.weight = weight
}

func (bb *rogueBarbell) getWeight() float64 {
	return bb.weight
}

func (bb *rogueBarbell) represent() string {
	return fmt.Sprintf(
		"This is a %v-kilo barbell for %v (%T)\n",
		bb.weight, bb.purpose, bb,
	)
}

type uralGiryaKettlebell struct {
	kettlebell
}

func (kb *uralGiryaKettlebell) setColor(color string) {
	kb.color = color
}

func (kb *uralGiryaKettlebell) setWeight(weight float64) {
	kb.weight = weight
}

func (kb *uralGiryaKettlebell) getColor() string {
	return kb.color
}

func (kb *uralGiryaKettlebell) getWeight() float64 {
	return kb.weight
}

func (kb *uralGiryaKettlebell) represent() string {
	return fmt.Sprintf(
		"This is a %v %v-kilo kettlebell (%T)\n",
		kb.color, kb.weight, kb,
	)
}

type uralGiryaBarbell struct {
	barbell
}

func (bb *uralGiryaBarbell) setPurpose(purpose string) {
	bb.purpose = purpose
}

func (bb *uralGiryaBarbell) setWeight(weight float64) {
	bb.weight = weight
}

func (bb *uralGiryaBarbell) getPurpose() string {
	return bb.purpose
}

func (bb *uralGiryaBarbell) getWeight() float64 {
	return bb.weight
}

func (bb *uralGiryaBarbell) represent() string {
	return fmt.Sprintf(
		"This is a %v-kilo barbell for %v (%T)\n",
		bb.weight, bb.purpose, bb,
	)
}
