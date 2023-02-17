package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

/*
Using bcrypt pkg for storing authentification w/o storing the actual password
*/

func main() {
	p := `NuggetInABiscuit123` //this is raw string => bytes

	hp, err := bcrypt.GenerateFromPassword([]byte(p), 13)
	/*@13 my poor old PC lost ability to immediately produce the output,
	so I stopped incrementing here. bcript.MaxCost would probably take a century
	*/
	if err != nil {
		fmt.Println("\nWe have a problem:\t", err)
	}
	fmt.Println(hp)
	fmt.Println(p)

	loginAttempt0 := `NuggetInABiscuit123`
	loginAttempt1 := `SafetyTorch999`

	//This is a successful login attempt
	err = bcrypt.CompareHashAndPassword([]byte(hp), []byte(loginAttempt0))
	if err != nil {
		fmt.Println("\nInvalid password")
		return
	}
	fmt.Println("Welcome back, Tobuscus")

	//This is a failed login attempt
	err = bcrypt.CompareHashAndPassword([]byte(hp), []byte(loginAttempt1))
	if err != nil {
		fmt.Println("\nInvalid password")
		return
	}
}
