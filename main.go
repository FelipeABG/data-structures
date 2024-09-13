package main

import "data-structures/structures"

func main() {
	q := structures.NewDynamicQueue[int]()

	q.Insert(5)
	q.Insert(6)
	q.Insert(7)

}
