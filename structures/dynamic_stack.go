package structures

import "errors"

type dynamicStack[T any] struct {
	data *linkedList[T]
}

func NewDynamicStack[T any]() *dynamicStack[T] {
	return &dynamicStack[T]{
		data: NewLinkedList[T](),
	}
}

func (s dynamicStack[T]) Empty() bool {
	return s.data.Empty()
}

func (s *dynamicStack[T]) Insert(element T) {
	s.data.AddLast(element)
}

func (s *dynamicStack[T]) Remove() (T, error) {

	if s.Empty() {
		return *new(T), errors.New("ERROR: Attempt to remove an element from an empty stack.")
	}

	result, _ := s.data.PopLast()

	return result, nil
}

func (s *dynamicStack[T]) Peek() (T, error) {

	if s.Empty() {
		return *new(T), errors.New("ERROR: Attempt to get an element from an empty stack.")
	}

}
