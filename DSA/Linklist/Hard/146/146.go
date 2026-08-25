package main

import "fmt"

// #TODO : Understand the problem more throughly
type LRUCache struct {
	cap    int
	head   *Node
	tail   *Node
	hasMap map[int]*Node
}

type Node struct {
	value, key int
	next       *Node
	prev       *Node
}

func Constructor(capacity int) LRUCache {
	h := &Node{value: -1}
	t := &Node{value: -1}

	h.next = t
	t.prev = h

	return LRUCache{cap: capacity, head: h, tail: t, hasMap: make(map[int]*Node, capacity)}
}

func (this *LRUCache) Get(key int) int {

}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.hasMap[key]; ok {
		node.value = value
		node.next.prev = node.prev
		node.prev.next = node.next

		this.head = node.next
		this.head.prev = node
		this.head.next = node.next.next
		return
	}
	newNode := &Node{value: value, key: key, prev: this.head, next: this.head.next}

}

func main() {
	fmt.Print("Hello")
}
