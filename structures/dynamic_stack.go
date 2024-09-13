package structures

import "errors"

type dynamicStack[T any] struct {
	data *linkedList[T]
}

// Creates and returns a reference to a dynamic stack
func NewDynamicStack[T any]() *dynamicStack[T] {
	return &dynamicStack[T]{
		data: NewLinkedList[T](),
	}
}

// Check if the stack is empty
func (s dynamicStack[T]) Empty() bool {
	return s.data.Empty()
}

// Inserts an element at the end of the stack
func (s *dynamicStack[T]) Insert(element T) {
	s.data.AddLast(element)
}

// Removes the last element of the stack
func (s *dynamicStack[T]) Remove() (T, error) {

	if s.Empty() {
		return *new(T), errors.New("ERROR: Attempt to remove an element from an empty stack.")
	}

	result, _ := s.data.PopLast()

	return result, nil
}

// Returns the last element of the stack
func (s *dynamicStack[T]) Peek() (T, error) {

	if s.Empty() {
		return *new(T), errors.New("ERROR: Attempt to get an element from an empty stack.")
	}

	target := s.data.first

	for target.next != nil {
		target = target.next
	}

	return target.value, nil
}
