package linked

import (
	"fmt"
	"os"
	"strconv"
)

// Data is an interface to allow storing arbitrary information inside an Elem
type Data interface{}

// Elem is a basic building block for a List
type Elem struct {
	ID   int
	Name string
	Data Data
	Next *Elem
}

// NewElem is a type Elem constructor
func NewElem(name string, data Data) *Elem {
	el := Elem{
		Name: name,
		Data: data,
	}
	if name[0] > '0' && name[0] < '9' {
		id, err := strconv.Atoi(name)
		if err == nil {
			el.ID = id
		}
	}

	return &el
}

func (e Elem) String() string {
	return fmt.Sprintf("<%v: %v>", e.Name, e.Data)
}

// List is an implementation of a linked list
type List struct {
	Name   string
	Head   *Elem
	Length int
}

// NewList is a type List constructor
func NewList(name string, head *Elem) *List {
	return &List{
		Name:   name,
		Head:   head,
		Length: 1,
	}
}

// Append adds a new last element to the list
func (l *List) Append(el *Elem) {
	var tail *Elem
	if l.Length == 1 {
		tail = l.Head
	} else {
		tail = l.penultimate().Next
	}
	tail.Next = el
	l.Length++
}

// Prepend adds a new first element to the list
func (l *List) Prepend(el *Elem) {
	el.Next = l.Head
	l.Head = el
	l.Length++
}

// Insert allows adding an element anywhere to the list
// with the name of the previous element provided
func (l *List) Insert(el *Elem, prevElName string) {
	prev, ok := l.SearchByName(prevElName)
	if !ok {
		err := fmt.Errorf("element with name '%v' not found", prevElName)
		_, _ = fmt.Fprint(os.Stderr, err)
		return
	}
	oldNext := prev.Next
	prev.Next = el
	el.Next = oldNext
	l.Length++
}

// PopHead removes and returns the first element in the list
func (l *List) PopHead() *Elem {
	oldHead := l.Head
	newHead := oldHead.Next
	l.Head = newHead
	l.Length--
	oldHead.Next = nil
	return oldHead
}

// PopTail removes and returns the last element in the list
func (l *List) PopTail() *Elem {
	penult := l.penultimate()
	tail := penult.Next
	penult.Next = nil // I have a bad feeling about this
	l.Length--
	return tail
}

// SearchByName supports element look up using Name field
func (l *List) SearchByName(elName string) (*Elem, bool) {
	if l.Head == nil {
		err := fmt.Errorf("empty or headless List")
		_, _ = fmt.Fprint(os.Stderr, err)
		return nil, false
	}
	cursor := l.Head

	for {
		if cursor.Name == elName {
			return cursor, true
		}
		if cursor.Next == nil {
			break
		}
		cursor = cursor.Next
	}
	err := fmt.Errorf("element with name '%v' not found in '%v'", elName, l.Name)
	_, _ = fmt.Fprint(os.Stderr, err)
	return nil, false
}

// Repr outputs a representation of a Linked List ina humanly-readable form
func (l *List) Repr() {
	fmt.Printf("%v:\n", l.Name)
	cursor := l.Head
	for {
		fmt.Printf("%v -->\t", cursor)
		if cursor.Next == nil {
			fmt.Print("<nil>\n")
			break
		}
		cursor = cursor.Next
	}
}

func (l *List) penultimate() *Elem {
	cursor := l.Head
	if l.Length == 1 {
		return cursor
	}

	for {
		if cursor.Next.Next == nil {
			return cursor
		}
		cursor = cursor.Next
	}
}
