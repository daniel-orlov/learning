package linked

import (
	"fmt"
	"os"
)

type Data interface{}

type Elem struct {
	ID   int
	Name string
	Data Data
	Next *Elem
}

func NewElem(name string, data Data) *Elem {
	el := Elem{
		Name: name,
		Data: data,
	}
	return &el
}

func (e Elem) String() string {
	return fmt.Sprintf("<%v: %v>", e.Name, e.Data)
}

type LinkedList struct {
	Name   string
	head   *Elem
	Length int
}

func NewLinkedList(name string, head *Elem) *LinkedList {
	return &LinkedList{
		Name:   name,
		head:   head,
		Length: 1,
	}
}

func (l *LinkedList) Append(el *Elem) {
	var tail *Elem
	if l.Length == 1 {
		tail = l.head
	} else {
		tail = l.penultimate().Next
	}
	tail.Next = el
	l.Length++
}

func (l *LinkedList) Prepend(el *Elem) {
	el.Next = l.head
	l.head = el
	l.Length++
}

func (l *LinkedList) Insert(el *Elem, prevElName string) {
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

func (l *LinkedList) PopHead() *Elem {
	oldHead := l.head
	newHead := oldHead.Next
	l.head = newHead
	l.Length--
	oldHead.Next = nil
	return oldHead
}

func (l *LinkedList) PopTail() *Elem {
	penult := l.penultimate()
	tail := penult.Next
	penult.Next = nil //I have a bad feeling about this
	l.Length--
	return tail
}

func (l *LinkedList) SearchByName(elName string) (*Elem, bool) {
	if l.head == nil {
		err := fmt.Errorf("empty or headless LinkedList")
		_, _ = fmt.Fprint(os.Stderr, err)
		return nil, false
	}
	cursor := l.head

	for {
		if cursor.Name == elName {
			return cursor, true
		}
		if cursor.Next == nil {
			break
		}
		cursor = cursor.Next
	}
	err := fmt.Errorf("element with name '%v' not found in '%v'.", elName, l.Name)
	_, _ = fmt.Fprint(os.Stderr, err)
	return nil, false
}

func (l *LinkedList) Repr() {
	fmt.Printf("%v:\n", l.Name)
	cursor := l.head
	for {
		fmt.Printf("%v -->\t", cursor)
		if cursor.Next == nil {
			fmt.Print("<nil>\n")
			break
		}
		cursor = cursor.Next
	}
}

func (l *LinkedList) penultimate() *Elem {
	cursor := l.head
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
