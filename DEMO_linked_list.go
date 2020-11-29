package main

import "fmt"

func main() {
	//Initiating the head & the list
	Elendil := NewElem("Elendil", nil)
	HofElendil := LinkedList{Name: "House of Elendil", head: Elendil, length: 1}

	//Appending an elem - a son of Elendil
	Isildur := NewElem("Isildur", nil)
	HofElendil.Append(Isildur)
	HofElendil.Repr()

	//Adding another son of Elendil
	Anarion := NewElem("Anarion", nil)
	HofElendil.Append(Anarion)
	HofElendil.Repr()

	//Wait a minute... We forgot to add Elros, the ancestor of Elendil.
	//Let's prepend him
	Elros := NewElem("Elros", nil)
	HofElendil.Prepend(Elros)
	HofElendil.Repr()

	//That's better! But what about Elrond?
	//Yeah, right...
	Elrond := NewElem("Elrond", nil)
	HofElendil.Prepend(Elrond)
	HofElendil.Repr()

	//Beor with me, fellow Tolkienists (especially the most attentive ones)
	//Let's quickly add some more
	Arvedui := NewElem("Arvedui", nil)
	HofElendil.Append(Arvedui)
	Aragorn := NewElem("Aragorn", nil)
	HofElendil.Append(Aragorn)

	//Let's see what we have for now, shall we?
	HofElendil.Repr()

	//Oops, we made quite a few mistakes :/
	//A day may come when we forsake this linked list and break all the pointers
	//but it is not this day.
	//So let's fix it!

	//First things first: we forgot Arathorn II, Aragorn's father
	//He should be between his son and Arvedui
	Arathorn := NewElem("Arathorn", nil)
	HofElendil.Insert(Arathorn, "Arvedui")

	//Then let's get rid of Elrond (sorry). He is a brother, not a father to Elros
	_ = HofElendil.PopHead()

	//Yes, I ignored him. Explicitly. Nothing personal.
	//Now let's see if he is still there
	if v, ok := HofElendil.SearchByName("Elrond"); ok {
		fmt.Println("We retrieved him: ", v)
	} else {
		fmt.Println(" Elrond sailed away with Galadriel")
	}

	//Are we done?
	HofElendil.Repr()

	//Oh, #@*&. We have Isildur's brother Anarion stuck there in the middle \-_- //
	//Well, before we decide to promote this to a feature, let's just say we fix this in Double Linked Lists
	//Thanks for joining me and have a great day!
	//
	//
	//
	//
	//
	//
	//Bye!
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//Come on, guys.
	//I have to stay late anyway and pick all the nil pointers that I dropped while implementing this.
}
