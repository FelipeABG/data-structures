package structures

import "errors"

type stack[T any] struct {
	top  int8
	size int8
	data []T
}

// Creates and return a reference to a stack
func NewStack[T any](size int8) *stack[T] {
	return &stack[T]{
		top:  -1,
		size: size,
		data: make([]T, size),
	}
}

// Checks if the stack is empty
func (s stack[T]) Empty() bool {
	return s.top == -1
}

// Checks if the stack is full
func (s *stack[T]) Full() bool {
	return s.top == s.size-1
}

// Inserts the element at the end of the stack
func (s *stack[T]) Insert(element T) error {

	if s.Full() {
		return errors.New("ERROR: Attempt to insert an element into a full stack")
	}

	s.top++
	s.data[s.top] = element

	return nil
}

// Removes the last element of the stack and returns it
func (s *stack[T]) Remove() (T, error) {

	if s.Empty() {
		return *new(T), errors.New("ERROR: Attempt to remove an element from an empty stack")
	}

	target := s.data[s.top]
	s.data[s.top] = *new(T)
	s.top--

	return target, nil
}

// Returns the last element of the stack.
func (s *stack[T]) Peek() (T, error) {
	if s.Empty() {
		return *new(T), errors.New("ERROR: Attempt to get an elemento from an empty stack")
	}

	return s.data[s.top], nil
}
