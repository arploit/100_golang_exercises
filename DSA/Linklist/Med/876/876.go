// https://leetcode.com/problems/middle-of-the-linked-list/description/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) insert(value int) {
	currentNode := head

	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	newNode := &ListNode{Val: value}
	currentNode.Next = newNode
}

func (head *ListNode) print() {
	currentNode := head

	for currentNode != nil {
		fmt.Printf("%d -->", currentNode.Val)
		currentNode = currentNode.Next
	}
	fmt.Println()
}

func middleNode(head *ListNode) *ListNode {
	length := 0
	currentNode := head

	for currentNode != nil {
		length += 1
		currentNode = currentNode.Next
	}
	if length < 2 {
		return currentNode
	}

	newHead := head
	for num := 0; num < length; num++ {
		if num == length/2 {
			return newHead
		}
		newHead = newHead.Next
	}

	return newHead
}

func reverseList(head *ListNode) *ListNode {
	currentNode := head
	var prev *ListNode = nil

	for currentNode != nil {
		next := currentNode.Next
		currentNode.Next = prev
		prev = currentNode
		currentNode = next
	}

	return prev
}

func main() {
	nodes := []int{1, 2, 3, 4, 5, 111, 22}
	if len(nodes) == 0 {
		return
	}
	newNode := &ListNode{Val: nodes[0]}
	for _, v := range nodes[1:] {
		newNode.insert(v)
	}
	newNode.print()
	newNode = reverseList(newNode)

	newNode.print()

}
