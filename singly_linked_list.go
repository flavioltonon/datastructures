package datastructures

import "fmt"

// SinglyLinkedList is a singly linked list of nodes carrying a value with a generic type.
type SinglyLinkedList[T any] struct {
	head   *Node[T]
	length int
}

// NewSinglyLinkedList creates an empty linked list of nodes containing objects with a type T.
func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return new(SinglyLinkedList[T])
}

// PrependNode inserts a given node at the beginning of the linked list. Time complexity: O(1).
func (l *SinglyLinkedList[T]) PrependNode(node *Node[T]) {
	l.length++

	if l.head == nil {
		l.head = node
		return
	}

	node.next = l.head
	l.head = node
}

// AppendNode inserts a given node at the end of the linked list. Time complexity: O(N).
func (l *SinglyLinkedList[T]) AppendNode(node *Node[T]) {
	if l.head == nil {
		l.head = node
		return
	}

	current := l.head

	for current.next != nil {
		current = current.next
	}

	current.next = node
	l.length++
}

// DeleteNode deletes a node with a given key. Time complexity: O(N).
func (l *SinglyLinkedList[T]) DeleteNode(key string) {
	if l.head == nil {
		return
	}

	if l.head.key == key {
		l.head = l.head.next
		l.length--
		return
	}

	current := l.head

	for {
		if current.next == nil {
			return
		}

		if current.next.key != key {
			current = current.next
			continue
		}

		current.next = current.next.next
		l.length--
	}
}

// Len returns the number of nodes in the linked list.
func (l *SinglyLinkedList[T]) Len() int {
	return l.length
}

// Print prints the nodes in the linked list. Time complexity: O(N).
func (l *SinglyLinkedList[T]) Print() {
	current := l.head

	for {
		if current != l.head {
			fmt.Printf(" -> ")
		}

		current.Print()

		if current == nil {
			fmt.Println()
			return
		}

		current = current.next
	}
}

// Search returns the first node identified by a given key. If there are no matches, it should return a nil
// node instead.
func (l *SinglyLinkedList[T]) Search(key string) *Node[T] {
	current := l.head

	for {
		if current == nil {
			return nil
		}

		if current.key == key {
			return current
		}

		current = current.next
	}
}
