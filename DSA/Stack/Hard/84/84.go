// Q: // https://leetcode.com/problems/largest-rectangle-in-histogram/description//
package main

import "fmt"

func largestRectangleArea(height []int) int {
	stack := []int{}
	area := 0

	for i := 0; i <= len(height); i++ {
		currentHeight := 0
		if i < len(height) {
			currentHeight = height[i]
		}
		{
			for len(stack) > 0 && currentHeight < height[stack[len(stack)-1]] {
				h := height[stack[len(stack)-1]]
				stack = stack[:len(stack)-1]
				width := 0
				if len(stack) == 0 {
					width = i
				} else {
					width = i - stack[len(stack)-1] - 1
				}
				tempArea := h * width
				if tempArea > area {
					area = tempArea
				}
			}
		}
		stack = append(stack, i)
	}

	return area
}

func main() {

	height := []int{1, 2, 3, 4, 2}
	fmt.Print("Output ", largestRectangleArea(height))

}
