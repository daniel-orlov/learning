package main

type Builder interface {
	setPropulsionType()
	setMaterialType()
	setCargoType()
	setName()
	getShip() *Ship
}

func getShipBuilder(builderType string) Builder {
	switch builderType {
	case "Galera":
		return newGaleraBuilder()
	case "FalconX":
		return newFalconXBuilder()
	default:
		return nil
	}
}
