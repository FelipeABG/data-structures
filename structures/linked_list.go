package structures

import "errors"

type node[T any] struct {
	value T
	next  *node[T]
}

type linkedList[T any] struct {
	first *node[T]
}

// Constructors
func newNode[T any](element T) *node[T] {
	return &node[T]{
		value: element,
		next:  nil,
	}
}

func NewLinkedList[T any]() *linkedList[T] {
	return &linkedList[T]{
		first: nil,
	}
}

// General purpouse methods
func (l linkedList[T]) Empty() bool {
	return l.first == nil
}

func (l linkedList[T]) Get(element T) (T, error) {

}

// Insertion methods
func (l *linkedList[T]) AddFirst(element T) {

	node := newNode(element)

	if l.Empty() {
		l.first = node
		return
	}

	node.next = l.first
	l.first = node

}

func (l *linkedList[T]) AddLast(element T) {
	node := newNode(element)

	if l.Empty() {
		l.first = node
		return
	}

	last := l.first
	for last.next != nil {
		last = last.next
	}

	last.next = node
}

func (l *linkedList[T]) AddSorted(element T) {

}

func (l *linkedList[T]) AddAfter(targer T, element T) {

}

// Deletion methods
func (l *linkedList[T]) PopFirst() (T, error) {

	if l.Empty() {
		return *new(T), errors.New("ERROR: Attempt to get an element from an empty linked list")
	}

	firstValue := l.first.value
	newFirst := l.first.next
	l.first = newFirst

	return firstValue, nil
}

func (l *linkedList[T]) PopLast() (T, error) {

	if l.Empty() {
		return *new(T), errors.New("ERROR: Attempt to get an element from an empty linked list")
	}

	if l.first.next == nil {
		lastValue := l.first.value
		l.first = nil
		return lastValue, nil
	}

	priorToLast := l.first
	var lastValue T

	for (priorToLast.next).next != nil {
		priorToLast = priorToLast.next
	}

	lastValue = priorToLast.next.value
	priorToLast.next = nil
	return lastValue, nil
}

func (l *linkedList[T]) Remove(element T) (T, error) {

}
