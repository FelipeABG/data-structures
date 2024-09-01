package structures

type node[T any] struct {
	value T
	next  *node[T]
}

type linkedList[T any] struct {
	first *node[T]
}

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

func (l linkedList[T]) Empty() bool {
	return l.first == nil
}

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
