package examples

import "github.com/flavioltonon/datastructures"

func ExampleSinglyLinkedList_Print() {
	linkedList := datastructures.NewSinglyLinkedList[int]()
	linkedList.AppendNode(datastructures.NewNode("b", 1))
	linkedList.AppendNode(datastructures.NewNode("c", 2))
	linkedList.PrependNode(datastructures.NewNode("a", 0))
	linkedList.Print()
	// Output: a: 0 -> b: 1 -> c: 2 -> <nil>
}
