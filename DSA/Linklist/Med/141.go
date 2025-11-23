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
		fmt.Println(currentNode.Val)
		currentNode = currentNode.Next
	}

}

//https://leetcode.com/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	currentNode := head
	var nodeMap = make(map[int]int)

	for currentNode.Next != nil {
		if _, ok := nodeMap[currentNode.Val]; !ok {
			nodeMap[currentNode.Val] = currentNode.Val
		} else {
			return true
		}
		currentNode = currentNode.Next
	}

	return false
}

//https://leetcode.com/problems/linked-list-cycle-ii
func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	prt1 := head
	prt2 := slow

	for prt1 != prt2 {
		prt1 = prt1.Next
		prt2 = prt2.Next

	}

	return prt1
}

func main() {
	nodes := []int{1, 2, 3, 4, 5, 111, 22}

	newNode := &ListNode{Val: nodes[0]}
	for _, v := range nodes[1:] {
		newNode.insert(v)
	}

	newNode.print()

}
