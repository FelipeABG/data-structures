package structures

import "errors"

type stack[T any] struct {
	top  int8
	size int8
	data []T
}

func NewStack[T any](size int8) *stack[T] {
	return &stack[T]{
		top:  -1,
		size: size,
		data: make([]T, size),
	}
}

func (s stack[T]) Empty() bool {
	return s.top == -1
}

func (s *stack[T]) Full() bool {
	return s.top == s.size-1
}

func (s *stack[T]) Insert(element T) error {

	if s.Full() {
		return errors.New("ERROR: Attempt to insert an element into a full stack")
	}

	s.top++
	s.data[s.top] = element

	return nil
}

func (s *stack[T]) Remove() (T, error) {

	if s.Empty() {
		return *new(T), errors.New("ERROR: Attempt to remove an element from an empty stack")
	}

	target := s.data[s.top]
	s.data[s.top] = *new(T)
	s.top--

	return target, nil
}

func (s stack[T]) Peek() T {

}
