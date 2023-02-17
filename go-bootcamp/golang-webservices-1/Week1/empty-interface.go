package main

import (
	"fmt"
	"strconv"
)

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Not enough money")
	}
	w.Cash -= amount
	return nil
}

func (w *Wallet) String() string {
	return "This wallet has " + strconv.Itoa(w.Cash) + " money units"
}

func main() {
	myWallet := &Wallet{Cash: 120}
	fmt.Printf("Raw payment: %#v\n", myWallet)
	fmt.Printf("Payment method: %s\n", myWallet)
}
