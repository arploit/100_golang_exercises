//https://leetcode.com/problems/trapping-rain-water/description/?envType=problem-list-v2&envId=array

package main

import (
	"fmt"
	"math"
)

func trap(height []int) int {
	left, right, totalArea := 0, len(height)-1, 0
	leftMax, rightMax := height[left], height[right]

	if len(height) <= 0 {
		return 0
	}

	for left < right {
		if leftMax < rightMax {
			left += 1
			leftMax = int(math.Max(float64(leftMax), float64(height[left])))
			totalArea += leftMax - height[left]
		} else {
			right -= 1
			rightMax = int(math.Max(float64(rightMax), float64(height[right])))
			totalArea += rightMax - height[right]
		}

	}

	return totalArea

}

func main() {

	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	height_2 := []int{4, 2, 0, 3, 2, 5}

	fmt.Println(trap(height))
	fmt.Println(trap(height_2))

}
