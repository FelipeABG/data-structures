package structures

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

}

func (s *stack[T]) Full() bool {

}

func (s *stack[T]) Insert(element T) {

}

func (s *stack[T]) Remove() T {

}

func (s stack[T]) Peek() T {

}
