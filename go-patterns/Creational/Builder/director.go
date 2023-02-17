package main

type Director struct {
	builder Builder
}

func newDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) SetBuilder(b Builder) {
	d.builder = b
}

func (d *Director) BuildShip() *Ship {
	d.builder.setName()
	d.builder.setMaterialType()
	d.builder.setCargoType()
	d.builder.setPropulsionType()
	return d.builder.getShip()
}
