package stack

import (
	ll "github.com/daniel-orlov/go-data-structures/linked"
)

type Stack struct {
	Name string
	list *ll.LinkedList
}

func NewStack(name string, elements ...*ll.elem) *Stack {
	stack := &Stack{
		Name: name,
		list: ll.NewLinkedList("stack", elements[0]),
	}
	for _, i := range elements[1:] {
		stack.list.Append(i)
	}
	return stack
}

func (s *Stack) Push(elements ...*ll.Elem) {
	for _, el := range elements {
		s.list.Append(el)
	}
}

func (s *Stack) Pop() *ll.Elem {
	return s.list.PopTail()
}
