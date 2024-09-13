package structures

import "errors"

type dynamicQueue[T any] struct {
	data *linkedList[T]
}

// Creates and returns a reference to a dynamic queue
func NewDynamicQueue[T any]() *dynamicQueue[T] {
	return &dynamicQueue[T]{
		data: NewLinkedList[T](),
	}
}

// Checks if the queue is empty
func (q dynamicQueue[T]) Empty() bool {
	return q.data.Empty()
}

// Inserts an element at the end of the queue
func (q *dynamicQueue[T]) Insert(element T) {
	q.data.AddLast(element)
}

// Removes the first element of the queue and returns it
func (q *dynamicQueue[T]) Remove(element T) (T, error) {

	if q.Empty() {
		return *new(T), errors.New("ERROR: Attempt to remove an element from an empty queue")
	}

	result, _ := q.data.PopFirst()

	return result, nil
}

// Returns the first element of the queue
func (q *dynamicQueue[T]) Get(element T) (T, error) {

	if q.Empty() {
		return *new(T), errors.New("ERROR: Attempt to remove an element from an empty queue")
	}

	return q.data.first.value, nil
}
