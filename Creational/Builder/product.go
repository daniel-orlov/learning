package main

import "fmt"

type Ship struct {
	Name           string
	MaterialType   string
	CargoType      string
	PropulsionType string
}

func (s *Ship) Represent() string {
	return fmt.Sprintf(
		"This is a %v built of the best %v to carry %v with a help of %v.\n",
		s.Name, s.MaterialType, s.CargoType, s.PropulsionType,
	)
}
