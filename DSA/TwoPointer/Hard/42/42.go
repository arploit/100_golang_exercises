// Q : https://leetcode.com/problems/trapping-rain-water/description/
package main

import "fmt"

func trap(height []int) int {
	res := 0
	for i := 0; i < len(height); i++ {

		left := height[i]
		for j := 0; j < i; j++ {
			left = max(left, height[j])
		}

		right := height[i]
		for j := i + 1; j < len(height); j++ {
			right = max(right, height[j])
		}

		res += min(left, right) - height[i]
	}

	return res
}

func trap_2(height []int) int {
	left := 0
	right := len(height) - 1
	leftMax, rightMax, res := 0, 0, 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				res += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				res += rightMax - height[right]
			}

			right--
		}
	}

	return res
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println("waterStored", trap(height))
	fmt.Println("waterStored", trap_2(height))

}
