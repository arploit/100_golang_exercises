package main

import "fmt"

type Node struct {
	prev *Node
	data int
	next *Node
}

type DDList struct {
	head *Node
}

func (l *DDList) insert(value int) {
	var newNode *Node = &Node{data: value}
	if l.head == nil {
		l.head = newNode
		return
	}

	var currentNode *Node = l.head

	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	currentNode.next = newNode
	newNode.prev = currentNode

}

func (l *DDList) print() {
	currentNode := l.head

	for currentNode != nil {
		fmt.Printf("%d -->", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println()
}

func (l *DDList) deleteNode(value int) {
	currentNode := l.head

	for currentNode != nil {
		if currentNode.data == value {
			if currentNode.next == nil {

				if currentNode.prev == nil {
					l.head = nil
					fmt.Println("Doubly linked list is deleted")
					return
				} else {
					currentNode.prev.next = nil
					fmt.Println("Last Element Deleted")
					return
				}

			} else {
				prevNode := currentNode.prev
				prevNode.next = currentNode.next
				currentNode.next.prev = prevNode
				return
			}
		}
		currentNode = currentNode.next
	}

}

func (l *DDList) reverseList() {
	currentNode := l.head
	var prev *Node = nil

	for currentNode != nil {

		nextNode := currentNode.next
		currentNode.next = prev
		currentNode.prev = nextNode

		prev = currentNode
		currentNode = nextNode
	}

	l.head = prev

}

func main() {

	newNode := &Node{data: 12}
	DDList := &DDList{head: newNode}

	DDList.insert(20)
	DDList.insert(30)
	DDList.insert(40)
	DDList.insert(50)
	DDList.print()
	DDList.reverseList()
	DDList.print()
}
