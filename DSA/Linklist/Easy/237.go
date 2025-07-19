package main

import "fmt"

func deleteNode(node *Node) {
	current := node
	if current.next != nil {
		current.data = current.next.data
		current.next = current.next.next
	}
}

func (list *linkList) length() {

	if list.head == nil {
		fmt.Println("null")
		return
	}

	var current *Node = list.head
	var length int32 = 1

	for current.next != nil {
		length += 1
		current = current.next
	}

	fmt.Println("list length: ", length)
}
