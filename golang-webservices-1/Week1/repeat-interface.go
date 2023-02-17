package main

import (
	"fmt"
)

type Payer interface {
	Pay(int) error
}

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

func Buy(p Payer, price int) {
	err := p.Pay(price)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Thank you for using %T!\n\n", p)
}

func main() {
	myByWallet := &Wallet{Cash: 65}
	Buy(myByWallet, 34)
}
