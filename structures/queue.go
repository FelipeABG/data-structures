package structures

import (
	"errors"
	"fmt"
)

// NOTE: THIS IS AN IMPLEMENTATION OF A CIRCULAR QUEUE
type queue[T any] struct {
	head int8
	tail int8
	size int8
	data []T
}

// Creates and returns a reference to a queue
func NewQueue[T any](size int8) *queue[T] {
	return &queue[T]{
		head: -1,
		tail: -1,
		size: size,
		data: make([]T, size),
	}
}

// Checks if the queue is empty
func (q *queue[T]) Empty() bool {
	return q.head == -1
}

// Checks if the queue is full
func (q *queue[T]) Full() bool {
	return (q.head == 0 && q.tail == q.size-1) || (q.tail+1 == q.head)
}

// Inserts an element into the end of the queue.
func (q *queue[T]) Insert(element T) error {
	if q.Full() {
		return errors.New("ERROR: Attempt to insert an element into a full queue")
	}

	if q.head == -1 {
		q.head = 0
	}

	// reset the tail "pointer" if the tail is at the last index
	if q.tail == q.size-1 {
		q.tail = -1
	}

	q.tail += 1
	q.data[q.tail] = element

	return nil
}

// Removes the first element of the queue and returns it
func (q *queue[T]) Remove() (error, T) {
	if q.Empty() {
		return errors.New("ERROR: Attempt to remove an element from an empty queue"), *new(T)
	}

	element := q.data[q.head]
	q.data[q.head] = *new(T)

	// Resets the whole queue if there is only one element
	if q.head == q.tail {
		q.head = -1
		q.tail = -1
	} else {
		q.head += 1
	}

	return nil, element
}

// Creates and returns a string with all queue elements
func (q *queue[T]) String() string {
	if q.Empty() {
		return ""
	}

	var result string
	i := q.head

	for {
		result += fmt.Sprintf("%v ", q.data[i])

		if i == q.tail {
			break
		}

		if i == q.size-1 {
			i = -1
		}

		i += 1
	}

	return result
}

// Return the firstt element of the queue.
func (q *queue[T]) Get() (error, T) {
	if q.Empty() {
		return errors.New("ERROR: Attempt to get an element from an empty queue"), *new(T)
	}

	return nil, q.data[q.head]
}
