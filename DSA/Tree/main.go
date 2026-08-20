package main

import (
	"fmt"
	tree104 "golang-exercise/DSA/Tree/Easy/104"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insert(values []interface{}) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: values[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for i < len(values) {
		node := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(values) && values[i] != nil {
			node.Left = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		// Right child
		if i < len(values) && values[i] != nil {
			node.Right = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func (root *TreeNode) insertOne(val int) {
	if val < root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{Val: val}
		} else {
			root.Left.insertOne(val)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Val: val}
		} else {
			root.Right.insertOne(val)
		}
	}
}

func (root *TreeNode) printLevelOrder() {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Print(node.Val, " ")

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

func convert(n *TreeNode) *tree104.TreeNode {
	if n == nil {
		return nil
	}
	return &tree104.TreeNode{
		Val:   n.Val,
		Left:  convert(n.Left),
		Right: convert(n.Right),
	}
}

func (root *TreeNode) Inorder() { //left Root Right
	if root == nil {
		return
	}

	root.Left.Inorder()
	fmt.Println(root.Val)
	root.Right.Inorder()
}

func (root *TreeNode) PreOrder() { //Root Left Right
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	root.Left.PreOrder()
	root.Right.PreOrder()
}

func (root *TreeNode) PostOrder() { //left Right root
	if root == nil {
		return
	}

	root.Left.PostOrder()
	root.Right.PostOrder()
	println(root.Val)

}

func isBalanced(root *TreeNode) bool {

	var height func(node *TreeNode) int

	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftHeight := height(node.Left)
		if leftHeight == -1 {
			return -1
		}

		rightHeight := height(node.Right)
		if rightHeight == -1 {
			return -1
		}

		if math.Abs(float64(leftHeight-rightHeight)) > 1 {
			return -1
		}
		return 1 + int(math.Max(float64(leftHeight), float64(rightHeight)))

	}
	return height(root) != -1

}

func main() {
	nums := []interface{}{3, 9, 20, nil, nil, 15, 7}
	root := insert(nums)
	root.printLevelOrder()
	isBalanced(root)
	// fmt.Println("")
	// root.Inorder()

	// fmt.Println("")
	// root.PreOrder()
	// fmt.Println("")
	// root.PostOrder()
	// fmt.Println("")
	fmt.Print("Output ", isBalanced(root))
	// tree.InvertTree(convert(root))
	// root.printLevelOrder()

}
