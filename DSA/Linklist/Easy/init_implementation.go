package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type linkList struct {
	head *Node
}

func (list *linkList) insert(data int) {
	newNode := &Node{data: data, next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head

	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func (list *linkList) print() {
	if list.head == nil {
		fmt.Println("Empty Linked list")
		return
	}

	var current *Node = list.head

	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println()
}

func main() {
	fmt.Println("Hello World!")

	myList := linkList{}

	myList.print()

	myList.insert(10)
	myList.insert(20)
	myList.insert(30)

	myList.print()

}
