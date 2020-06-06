//package Dog implements functions to covert age between dogs and humans
package dog

//DogYears converts human age into dog age
func DogYears(humanAge int) int {
	return int(humanAge / 7)
}

//HumanYears converts dog age into human age
func HumanYears(dogAge int) int {
	return dogAge * 7
}