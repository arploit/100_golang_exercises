package main

import (
	"fmt"
)

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

// https://leetcode.com/problems/linked-list-cycle/
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

// https://leetcode.com/problems/linked-list-cycle-ii
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

func oddEvenList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	odd := head
	even := head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = odd.Next.Next
		even.Next = even.Next.Next

		odd = odd.Next
		even = even.Next
	}

	odd.Next = evenHead

	return head

}

func createLoop(head *ListNode, pos int) {
	if pos < 0 {
		return
	}

	current := head
	var loopNode *ListNode
	index := 0

	// Traverse to find loopNode at given pos
	for current != nil {
		if index == pos {
			loopNode = current
		}
		if current.Next == nil {
			break // last node
		}
		current = current.Next
		index++
	}

	// Connect last node → loopNode
	if loopNode != nil {
		current.Next = loopNode
	}
}

// https://takeuforward.org/linked-list/length-of-loop-in-linked-list
func (head *ListNode) findLengthOfLoop() int {
	slow := head
	fast := head
	length := 1

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		fmt.Println("fast", fast.Val)
		slow = slow.Next
		if slow == fast {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return 0
	}

	p1 := head
	p2 := slow

	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}

	lengthP := p1.Next

	for p1 != lengthP {
		length += 1
		lengthP = lengthP.Next
	}

	return length
}

// https://leetcode.com/problems/palindrome-linked-list/
// 234. Palindrome Linked List
func isPalindrome(head *ListNode) bool {

	var stack = []int{}

	currentNode := head

	for currentNode != nil {
		stack = append(stack, currentNode.Val)
		currentNode = currentNode.Next
	}

	currentNode = head
	for currentNode != nil {
		if currentNode.Val != stack[len(stack)-1] {
			return false
		}
		stack = stack[:len(stack)-1]
		currentNode = currentNode.Next
	}

	return true

}

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	currentNode := head
	for currentNode != nil {
		currentNode = currentNode.Next
		length += 1
	}

	if n == length {
		return head.Next
	}
	newHead := head
	for i := 0; i < length-n-1; i++ {
		newHead = newHead.Next

	}

	newHead.Next = newHead.Next.Next

	return head
}

// 2095: https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/
func deleteMiddle(head *ListNode) *ListNode {
	currentNode := head
	var length int = 0

	for currentNode != nil {
		length += 1
		currentNode = currentNode.Next
	}
	if length == 1 {
		return nil
	}

	newHead := head
	for i := 0; i < (length/2)-1; i++ {
		newHead = newHead.Next
	}

	if newHead.Next == nil {
		newHead.Next = nil
		return head
	}

	newHead.Next = newHead.Next.Next
	return head

}

func findMiddle(head *ListNode) *ListNode {

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func merge(left *ListNode, right *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for left != nil && right != nil {
		if left.Val <= right.Val {
			tail.Next = left
			left = left.Next
		} else {
			tail.Next = right
			right = right.Next
		}

		tail = tail.Next

	}
	if left != nil {
		tail.Next = left
	} else {
		tail.Next = right
	}

	return dummy.Next
}

// 148 : https://leetcode.com/problems/sort-list/description/
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	middle := findMiddle(head)
	right := middle.Next
	middle.Next = nil

	left := sortList(head)
	right = sortList(right)

	return merge(left, right)

}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]bool)
	tempA := headA
	tempB := headB

	for tempA != nil {
		m[tempA] = true
		tempA = tempA.Next
	}

	for tempB != nil {
		if m[tempB] {
			return tempB
		} else {
			tempB = tempB.Next
		}
	}

	return nil

}

func (head *ListNode) length() int {
	currentHead := head
	l := 0

	for currentHead != nil {
		l += 1
		currentHead = currentHead.Next
	}

	return l
}

func openAWindow(head *ListNode, windowSize int) *ListNode {
	currentHead := head
	for currentHead.Next != nil || windowSize > 0 {
		currentHead = currentHead.Next
		windowSize -= 1
	}

	return currentHead
}

func swap(A *ListNode, B *ListNode) (*ListNode, *ListNode) {
	tempA := A
	A.Next = B.Next
	B.Next = tempA.Next
	return B, A
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	length := head.length()
	tempA := head
	tempB := head

	for length >= k {
		if tempB.Next != nil {
			tempB = openAWindow(tempB, k)
			tempA, tempB = swap(tempA, tempB)
			length = length - k
			tempA = tempB.Next
			tempB = tempB.Next
		} else {
			return head
		}

	}
	return head

}

func main() {
	nodes := []int{1, 2, 3, 4, 5}
	k := 2

	newNode := &ListNode{Val: nodes[0]}
	for _, v := range nodes[1:] {
		newNode.insert(v)
	}

	reverseKGroup(newNode, k).print()

	// deleteMiddle(newNode).print()

	// fmt.Println("isPalindrone ", isPalindrome(newNode))

	// newNode.print()

	// oddEvenNod := oddEvenList(newNode)

	// oddEvenNod.print()

	// oddEvenNode := oddEvenList(newNode)

	// createLoop(newNode, 1)

	// if hasCycle(newNode) {
	// 	fmt.Println("Cycle detected!")
	// }

	// entry := detectCycle(newNode)
	// fmt.Println("Cycle starts at:", entry.Val)

	// fmt.Print("findLengthOfLoop ", newNode.findLengthOfLoop())

	// oddEvenNode.print()

	// newNode.print()

}
