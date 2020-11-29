package main

import "fmt"

var hand = []*elem{
	NewElem("Jack of Clubs", "Adam Wilber"),
	NewElem("Queen of Clubs", "Laura London"),
	NewElem("King of Hearts", "Peter McKinnon"),
	NewElem("Jack of Spades", "Gianni IIRC"),
}

func main() {
	//Creating a deck of cards
	KoS := NewElem("King of Spades", "Daniel Madison")
	deckOfCards := NewStack("Rounders", KoS)

	//Adding a hand-full of cards to a stack
	deckOfCards.Push(hand...)

	//Popping them one by one
	for deckOfCards.list.length > 1 {
		fmt.Println(deckOfCards.Pop())
	}
}
