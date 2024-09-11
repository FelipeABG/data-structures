package structures

import "errors"

type dynamicQueue[T any] struct {
	data *linkedList[T]
}

func NewDynamicQueue[T any]() *dynamicQueue[T] {
	return &dynamicQueue[T]{
		data: NewLinkedList[T](),
	}
}

func (q dynamicQueue[T]) Empty() bool {
	return q.data.Empty()
}

func (q *dynamicQueue[T]) Insert(element T) {
	q.data.AddLast(element)
}

func (q *dynamicQueue[T]) Remove(element T) (T, error) {

	if q.Empty() {
		return *new(T), errors.New("ERROR: Attempt to remove an element from an empty queue")
	}

	result, _ := q.data.PopFirst()

	return result, nil
}

func (q *dynamicQueue[T]) Get(element T) (T, error) {

	if q.Empty() {
		return *new(T), errors.New("ERROR: Attempt to remove an element from an empty queue")
	}

	return q.data.first.value, nil
}
