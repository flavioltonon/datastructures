package datastructures

import "fmt"

// Node is a datastructure node that is identified by a key and carries a generic value.
type Node[T any] struct {
	key   string
	value T
	next  *Node[T]
}

// NewNode creates a new Node.
func NewNode[T any](key string, value T) *Node[T] {
	return &Node[T]{
		key:   key,
		value: value,
	}
}

// Print prints the key/value pair of the Node.
func (n *Node[T]) Print() {
	if n == nil {
		fmt.Print(n)
		return
	}

	fmt.Printf("%s: %+v", n.key, n.value)
}
