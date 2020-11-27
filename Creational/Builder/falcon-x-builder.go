package main

type FalconXBuilder struct {
	Name           string
	MaterialType   string
	CargoType      string
	PropulsionType string
}

func newFalconXBuilder() *FalconXBuilder {
	return &FalconXBuilder{}
}

func (b *FalconXBuilder) setPropulsionType() {
	b.PropulsionType = "Rocket Engine"
}

func (b *FalconXBuilder) setMaterialType() {
	b.MaterialType = "Enchanted Steel"
}

func (b *FalconXBuilder) setCargoType() {
	b.CargoType = "Astronauts and Memes"
}

func (b *FalconXBuilder) setName() {
	b.Name = "FalconX Disco Dino Edition"
}

func (b *FalconXBuilder) getShip() *Ship {
	return &Ship{
		Name:           b.Name,
		MaterialType:   b.MaterialType,
		CargoType:      b.CargoType,
		PropulsionType: b.PropulsionType,
	}
}
