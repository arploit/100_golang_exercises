// Q :  https://leetcode.com/problems/container-with-most-water/

package main

import "fmt"

func maxArea(height []int) int {
	area := 0
	if len(height) <= 1 {
		return 0
	}

	left := 0
	right := len(height) - 1

	for left < right {
		tempArea := 0
		if height[left] >= height[right] {
			tempArea = height[right] * (right - left)
			right -= 1
		} else {
			tempArea = height[left] * (right - left)
			left += 1
		}
		if tempArea > area {
			area = tempArea
		}
	}

	return area
}

func main() {
	height := []int{1, 1}
	fmt.Println(maxArea(height))
}
