package stack

import (
	ll "github.com/daniel-orlov/go-data-structures/linked"
)

// Stack is a basic implementation of a LIFO stack using LinkedList
type Stack struct {
	// Name is an ID field for a Stack
	Name string
	// List is an actual linked list behind the Stack
	List *ll.LinkedList
}

// NewStack is a type Stack consctructor
func NewStack(name string, elements ...*ll.Elem) *Stack {
	stack := &Stack{
		Name: name,
		List: ll.NewLinkedList("stack", elements[0]),
	}
	for _, i := range elements[1:] {
		stack.List.Append(i)
	}
	return stack
}

// Push supports adding 1 or more Elems to the stack
func (s *Stack) Push(elements ...*ll.Elem) {
	for _, el := range elements {
		s.List.Append(el)
	}
}

// Pop removes the topmost Elem from the Stack and returns it
func (s *Stack) Pop() *ll.Elem {
	return s.List.PopTail()
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.List.Length > 0
}
