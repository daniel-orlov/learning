package main

import "fmt"

type Phone struct {
	Money   int
	PhoneID string
}

func (ph *Phone) Pay(amount int) error {
	if ph.Money < amount {
		return fmt.Errorf("Not enough money")
	}
	ph.Money -= amount
	return nil
}

func (ph *Phone) Ring(number string) error {
	if number == "" {
		return fmt.Errorf("Please, enter the phone")
	}
	return nil
}

//-------------------------------------

type Payer interface {
	Pay(int) error
}

type Ringer interface {
	Ring(string) error
}

type NFCPhone interface {
	Payer
	Ringer
}

//-------------------------------------

func PayForTunnelbanaWithPhone(phone NFCPhone) {
	err := phone.Pay(1)
	if err != nil {
		fmt.Printf("Transaction failed %v\n\n", err)
		return
	}
	fmt.Printf("You paid with %T. Have a pleasant trip!", phone)
}

//-------------------------------------

func main() {
	myPhone := &Phone{Money: 9}
	PayForTunnelbanaWithPhone(myPhone)
}
