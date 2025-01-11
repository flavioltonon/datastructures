package main

import "github.com/flavioltonon/datastructures"

func main() {
	linkedList := datastructures.NewSinglyLinkedList[int]()
	linkedList.AppendNode(datastructures.NewNode("b", 1))
	linkedList.AppendNode(datastructures.NewNode("c", 2))
	linkedList.AppendNode(datastructures.NewNode("d", 3))
	linkedList.PrependNode(datastructures.NewNode("a", 0))
	linkedList.Traverse()
}
