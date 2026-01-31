// Q:// https://leetcode.com/problems/reverse-nodes-in-k-group/description/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getLength(head *ListNode) int {
	l := 0
	for head != nil {
		l++
		head = head.Next
	}
	return l
}
func reverseK(head *ListNode, k int) (*ListNode, *ListNode) {
	var prev *ListNode
	curr := head

	for k > 0 {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
		k--
	}

	return prev, head
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}

	length := getLength(head)
	dummy := &ListNode{Next: head}
	prevGroupTail := dummy
	curr := head

	for length >= k {

		nextGroupHead := curr
		for i := 0; i < k; i++ {
			nextGroupHead = nextGroupHead.Next
		}

		newHead, newTail := reverseK(curr, k)
		prevGroupTail.Next = newHead
		newTail.Next = nextGroupHead

		prevGroupTail = newTail
		curr = nextGroupHead
		length -= k
	}

	return dummy.Next
}

func main() {

}
