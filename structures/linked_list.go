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

// General purpose methods
func (l linkedList[T]) Empty() bool {
	return l.first == nil
}

func (l linkedList[T]) Get(element T) (T, error) {

	if l.Empty() {
		return *new(T), errors.New("ERROR: Attempt to get an element from an empty linked list")
	}

	target := l.first

	for target != nil {
		if target.value == element {
			return target.value, nil
		}
		target = target.next
	}

	return *new(T), errors.New("ERROR: Attempt to get a non-existent element from the linked list")

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
		return *new(T), errors.New("ERROR: Attempt to remove an element from an empty linked list")
	}

	firstValue := l.first.value
	newFirst := l.first.next
	l.first = newFirst

	return firstValue, nil
}

func (l *linkedList[T]) PopLast() (T, error) {

	if l.Empty() {
		return *new(T), errors.New("ERROR: Attempt to remove an element from an empty linked list")
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

func (l *linkedList[T]) Pop(element T) error {

	if l.Empty() {
		return errors.New("ERROR: Attempt to remove an element from an empty linked list")
	}

	if l.first.value == element {
		l.first = l.first.next
		return nil
	}

	prior := l.first

	for prior.next != nil {
		if prior.next.value == element {
			prior.next = prior.next.next
			return nil
		}
		prior = prior.next
	}

	return errors.New("ERROR: Attempt to pop a non-existent element from the linked list")
}
