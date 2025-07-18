package main

func deleteNode(node *Node) {
	current := node
	if current.next != nil {
		current.data = current.next.data
		current.next = current.next.next
	}
}
