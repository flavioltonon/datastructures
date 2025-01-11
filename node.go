package datastructures

import "fmt"

type Node[T any] struct {
	key   string
	value T
	next  *Node[T]
}

func NewNode[T any](key string, value T) *Node[T] {
	return &Node[T]{
		key:   key,
		value: value,
	}
}

func (n *Node[T]) Print() {
	if n == nil {
		fmt.Print(n)
		return
	}

	fmt.Printf("%s: %+v", n.key, n.value)
}
