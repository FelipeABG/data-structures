package structures

type node[T any] struct {
	value T
	next  *node[T]
}

type linkedList[T any] struct {
	first *node[T]
}
