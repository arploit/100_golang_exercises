package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0

	var height func(node *TreeNode) int

	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftHeight := height(node.Left)
		rightHeight := height(node.Right)

		maxDiameter = int(math.Max(float64(maxDiameter), float64(leftHeight+rightHeight)))

		return 1 + int(math.Max(float64(leftHeight), float64(rightHeight)))

	}

	height(root)
	return maxDiameter
}
