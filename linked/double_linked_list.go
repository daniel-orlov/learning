package linked

import (
	"fmt"
	"os"
)

// Element is a double-linked list specific type for
type Element struct {
	name string
	data Data
	next *Element
	prev *Element
}

// NewElement is a type Element constructor
func NewElement(name string, data Data) *Element {
	el := Element{
		name: name,
		data: data,
	}
	return &el
}

func (e Element) String() string {
	return fmt.Sprintf("<%v: %v>", e.name, e.data)
}

// DoubleList is an implementation of a double-linked list
type DoubleList struct {
	Name   string
	head   *Element
	tail   *Element
	length int
}

// NewDoubleList is a type DoubleList constructor
func NewDoubleList(name string, head *Element) *DoubleList {
	return &DoubleList{
		Name:   name,
		head:   head,
		tail:   head,
		length: 1,
	}
}

// Append adds a new last element to the list
func (l *DoubleList) Append(el *Element) {
	l.tail.next = el
	el.prev = l.tail
	l.tail = el
	l.length++
}

// Prepend adds a new first element to the list
func (l *DoubleList) Prepend(el *Element) {
	el.next = l.head
	l.head.prev = el
	l.head = el
	l.length++
}

// InsertAfterElement allows adding an Element after any other one on a list
// It supports searching both head-to-tail and backwards
func (l *DoubleList) InsertAfterElement(el *Element, elName string, backwards bool) {
	prev, ok := l.SearchByName(elName, backwards)
	if !ok {
		err := fmt.Errorf("Element with name '%v' not found", elName)
		_, _ = fmt.Fprint(os.Stderr, err)
		return
	}
	oldNext := prev.next
	prev.next = el
	el.prev = prev
	el.next = oldNext
	oldNext.prev = el
	l.length++
}

// InsertBeforeElement allows adding an Element before any other one on a list
// It supports searching both head-to-tail and backwards
func (l *DoubleList) InsertBeforeElement(el *Element, elName string, backwards bool) {
	next, ok := l.SearchByName(elName, backwards)
	if !ok {
		err := fmt.Errorf("Element with name '%v' not found", elName)
		_, _ = fmt.Fprint(os.Stderr, err)
		return
	}
	oldPrev := next.prev
	next.prev = el
	el.next = next
	el.prev = oldPrev
	oldPrev.next = el
	l.length++
}

// PopHead removes the first element from the list and returns it
func (l *DoubleList) PopHead() *Element {
	oldHead := l.head
	newHead := oldHead.next
	l.head = newHead
	oldHead.next = nil
	newHead.prev = nil
	l.length--
	return oldHead
}

// PopTail removes the last element from the list and returns it
func (l *DoubleList) PopTail() *Element {
	oldTail := l.tail
	newTail := l.tail.prev
	l.tail = newTail
	oldTail.prev = nil
	newTail.next = nil
	l.length--
	return oldTail
}

// SearchByName supports element look up using Name field (both head-to-tail and backwards)
func (l *DoubleList) SearchByName(elName string, backwards bool) (*Element, bool) {
	if l.head == nil {
		err := fmt.Errorf("empty DoubleList")
		_, _ = fmt.Fprint(os.Stderr, err)
		return nil, false
	}
	if !backwards {
		cursor := l.head

		for {
			if cursor.name == elName {
				return cursor, true
			}
			if cursor.next == nil {
				break
			}
			cursor = cursor.next
		}
	} else {
		cursor := l.tail

		for {
			if cursor.name == elName {
				return cursor, true
			}
			if cursor.prev == nil {
				break
			}
			cursor = cursor.prev
		}
	}

	// not found
	err := fmt.Errorf("element with name '%v' not found in '%v'.", elName, l.Name)
	_, _ = fmt.Fprint(os.Stderr, err)
	return nil, false
}

// Repr outputs a representation of a Linked List ina humanly-readable form
func (l *DoubleList) Repr(backwards bool) {
	fmt.Printf("%v:\n", l.Name)
	if !backwards {
		fmt.Println("Head to tail: ")
		cursor := l.head
		for {
			fmt.Printf("%v <-->\t", cursor)
			if cursor.next == nil {
				fmt.Print("<nil>\n")
				break
			}
			cursor = cursor.next
		}
		return
	}
	fmt.Println("Tail to head: ")
	cursor := l.tail
	for {
		fmt.Printf("%v <-->\t", cursor)
		if cursor.prev == nil {
			fmt.Print("<nil>\n")
			break
		}
		cursor = cursor.prev
	}
	return
}
