package linked

import (
	"fmt"
	"os"
)

type Data interface{}

type Elem struct {
	name string
	data Data
	next *Elem
	//prev *Elem //will use for Double LLs
}

func NewElem(name string, data Data) *Elem {
	el := Elem{
		name: name,
		data: data,
	}
	return &el
}

func (e Elem) String() string {
	return fmt.Sprintf("<%v: %v>", e.name, e.data)
}

type LinkedList struct {
	Name   string
	head   *Elem
	length int
}

func NewLinkedList(name string, head *Elem) *LinkedList {
	return &LinkedList{
		Name:   name,
		head:   head,
		length: 1,
	}
}

func (l *LinkedList) Append(el *Elem) {
	var tail *Elem
	if l.length == 1 {
		tail = l.head
	} else {
		tail = l.penultimate().next
	}
	tail.next = el
	l.length++
}

func (l *LinkedList) Prepend(el *Elem) {
	el.next = l.head
	l.head = el
	l.length++
}

func (l *LinkedList) Insert(el *Elem, prevElName string) {
	prev, ok := l.SearchByName(prevElName)
	if !ok {
		err := fmt.Errorf("element with name '%v' not found", prevElName)
		_, _ = fmt.Fprint(os.Stderr, err)
		return
	}
	oldNext := prev.next
	prev.next = el
	el.next = oldNext
	l.length++
}

func (l *LinkedList) PopHead() *Elem {
	oldHead := l.head
	newHead := oldHead.next
	l.head = newHead
	l.length--
	oldHead.next = nil
	return oldHead
}

func (l *LinkedList) PopTail() *Elem {
	penult := l.penultimate()
	tail := penult.next
	penult.next = nil //I have a bad feeling about this
	l.length--
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
		if cursor.name == elName {
			return cursor, true
		}
		if cursor.next == nil {
			break
		}
		cursor = cursor.next
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
		if cursor.next == nil {
			fmt.Print("<nil>\n")
			break
		}
		cursor = cursor.next
	}
}

func (l *LinkedList) penultimate() *Elem {
	cursor := l.head
	if l.length == 1 {
		return cursor
	}

	for {
		if cursor.next.next == nil {
			return cursor
		}
		cursor = cursor.next
	}
}
