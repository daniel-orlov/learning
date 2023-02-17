package main

func main() {
	client := &Client{}
	laptopWithWindows := &windows{}

	client.runOS(laptopWithWindows)

	linuxMachine := &linux{}
	linuxVMOnWindows := &vmAdapter{
		linuxImage: linuxMachine,
	}

	client.runOS(linuxVMOnWindows)
}
