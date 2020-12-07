package main

type caramelLayer struct {
	ketoCake KetoCake
}

func (c *caramelLayer) getPrice() int {
	cakePrice := c.ketoCake.getPrice()
	return cakePrice + 20
}

type chocolateDecorations struct {
	ketoCake KetoCake
}

func (c *chocolateDecorations) getPrice() int {
	cakePrice := c.ketoCake.getPrice()
	return cakePrice + 15
}
