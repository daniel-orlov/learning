package main

func main() {
	//Creating printers
	hpPrinter := &hp{}
	epsonPrinter := &epson{}

	//Creating a device - my android phone
	onePlusPlus := &androidPhone{}

	//Connecting a printer and printing a test file
	onePlusPlus.setPrinter(hpPrinter)
	onePlusPlus.print()

	//Creating a device - my linux laptop
	mcAcer := &linuxLaptop{}

	//Connecting a printer and printing a test file
	mcAcer.setPrinter(epsonPrinter)
	mcAcer.print()

	//And now swapping them
	onePlusPlus.setPrinter(epsonPrinter)
	mcAcer.setPrinter(hpPrinter)
	onePlusPlus.print()
	mcAcer.print()
}
