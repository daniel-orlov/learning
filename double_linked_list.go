package main

import (
	"fmt"
	"os"
)

type element struct {
	name string
	data Data
	next *element
	prev *element
}

func NewElement(name string, data Data) *element {
	el := element{
		name: name,
		data: data,
	}
	return &el
}

func (e element) String() string {
	return fmt.Sprintf("<%v: %v>", e.name, e.data)
}

type DoubleLinkedList struct {
	Name   string
	head   *element
	tail   *element
	length int
}

func NewDoubleLinkedList(name string, head *element) *DoubleLinkedList {
	return &DoubleLinkedList{
		Name:   name,
		head:   head,
		tail:   head,
		length: 1,
	}
}

func (l *DoubleLinkedList) Append(el *element) {
	l.tail.next = el
	el.prev = l.tail
	l.tail = el
	l.length++
}

func (l *DoubleLinkedList) Prepend(el *element) {
	el.next = l.head
	l.head.prev = el
	l.head = el
	l.length++
}

func (l *DoubleLinkedList) InsertAfterElement(el *element, elName string, backwards bool) {
	prev, ok := l.SearchByName(elName, backwards)
	if !ok {
		err := fmt.Errorf("element with name '%v' not found", elName)
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

func (l *DoubleLinkedList) InsertBeforeElement(el *element, elName string, backwards bool) {
	next, ok := l.SearchByName(elName, backwards)
	if !ok {
		err := fmt.Errorf("element with name '%v' not found", elName)
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

func (l *DoubleLinkedList) PopHead() *element {
	oldHead := l.head
	newHead := oldHead.next
	l.head = newHead
	oldHead.next = nil
	newHead.prev = nil
	l.length--
	return oldHead
}

func (l *DoubleLinkedList) PopTail() *element {
	oldTail := l.tail
	newTail := l.tail.prev
	l.tail = newTail
	oldTail.prev = nil
	newTail.next = nil
	l.length--
	return oldTail
}

func (l *DoubleLinkedList) SearchByName(elName string, backwards bool) (*element, bool) {
	if l.head == nil {
		err := fmt.Errorf("empty DoubleLinkedList")
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
	//not found
	err := fmt.Errorf("element with name '%v' not found in '%v'.", elName, l.Name)
	_, _ = fmt.Fprint(os.Stderr, err)
	return nil, false
}

func (l *DoubleLinkedList) Repr(backwards bool) {
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
