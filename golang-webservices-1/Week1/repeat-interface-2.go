package main

import (
    "fmt"
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
//------------------------------------------
type Card struct {
    Balance int
    ValidThrough string
    CardholderName string
    CVV string
    Number string
}

func (c *Card) Pay(amount int) error {
    if c.Balance < amount {
        return fmt.Errorf("Not enough money")
    }
    c.Balance -= amount
    return nil
}
//------------------------------------------
type ApplePay struct {
    Money int
    AppleID string
}

func (a *ApplePay) Pay(amount int) error {
    if a.Money < amount {
        return fmt.Errorf("Not enough money")
    }
    a.Money -= amount
    return nil
}
//------------------------------------------

type Payer interface {
    Pay(int) error
}

func Buy(p Payer, price int) {
    err := p.Pay(price)
    if err != nil {
        fmt.Printf("Transaction failed when using %T: %v\n\n", p, err)
        return
    }
    fmt.Printf("Thank you for using %T!\n\n", p)
}

func main() {
    myWallet := &Wallet{Cash: 65}
    Buy(myWallet, 34)

    var myMoney Payer
    myMoney = &Card{Balance: 1500, CardholderName: "Bifyllk Hijhaked"}
    Buy(myMoney, 999)

    myMoney = &ApplePay{Money: 0}
    Buy(myMoney, 1)

}
