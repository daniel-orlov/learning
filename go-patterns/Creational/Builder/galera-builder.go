package main

type GaleraBuilder struct {
	Name           string
	MaterialType   string
	CargoType      string
	PropulsionType string
}

func newGaleraBuilder() *GaleraBuilder {
	return &GaleraBuilder{}
}

func (b *GaleraBuilder) setPropulsionType() {
	b.PropulsionType = "Oarsmen and Sails"
}

func (b *GaleraBuilder) setMaterialType() {
	b.MaterialType = "Lebanese Ceder"
}

func (b *GaleraBuilder) setCargoType() {
	b.CargoType = "Goods and IT-services"
}

func (b *GaleraBuilder) setName() {
	b.Name = "Typical Galera"
}

func (b *GaleraBuilder) getShip() *Ship {
	return &Ship{
		Name:           b.Name,
		MaterialType:   b.MaterialType,
		CargoType:      b.CargoType,
		PropulsionType: b.PropulsionType,
	}
}
