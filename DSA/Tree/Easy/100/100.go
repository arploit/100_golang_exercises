package main

import "fmt"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

// Q:// https://leetcode.com/problems/same-tree/

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == nil && q == nil
	}

	return p.value == q.value &&
		isSameTree(p.left, q.left) && isSameTree(p.right, q.right)

}

func main() {
	fmt.Printf("Tree")
}
