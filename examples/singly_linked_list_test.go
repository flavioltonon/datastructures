package examples

import "github.com/flavioltonon/datastructures"

func ExampleSinglyLinkedList() {
	linkedList := datastructures.NewSinglyLinkedList[int]()
	linkedList.AppendNode(datastructures.NewNode("b", 1))
	linkedList.AppendNode(datastructures.NewNode("c", 2))
	linkedList.AppendNode(datastructures.NewNode("d", 3))
	linkedList.PrependNode(datastructures.NewNode("a", 0))
	linkedList.Traverse()
	// Output: a: 0 -> b: 1 -> c: 2 -> d: 3 -> <nil>
}
