package main

import (
	"fmt"

	ll "github.com/daniel-orlov/go-data-structures/linked"
	"github.com/daniel-orlov/go-data-structures/stack"
)

var hand = []*ll.Elem{
	ll.NewElem("Jack of Clubs", "Adam Wilber"),
	ll.NewElem("Queen of Clubs", "Laura London"),
	ll.NewElem("King of Hearts", "Peter McKinnon"),
	ll.NewElem("Jack of Spades", "Gianni IIRC"),
}

func main() {
	//Creating a deck of cards
	KoS := ll.NewElem("King of Spades", "Daniel Madison")
	deckOfCards := stack.NewStack("Rounders", KoS)

	//Adding a hand-full of cards to a stack
	deckOfCards.Push(hand...)

	//Popping them one by one
	for deckOfCards.List.Length > 1 {
		fmt.Println(deckOfCards.Pop())
	}
}
