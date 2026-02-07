// Q: // https://leetcode.com/problems/search-in-rotated-sorted-array/description/

package main

import (
	"fmt"
)

func bs(nums []int, start int, end int, target int) int {
	if start > end {
		return -1
	}

	middle := (start + end) / 2

	if nums[middle] == target {
		return middle
	}

	if nums[start] <= nums[middle] {
		if target >= nums[start] && target < nums[middle] {
			return bs(nums, start, middle-1, target)
		}
		return bs(nums, middle+1, end, target)
	}

	if target > nums[middle] && target <= nums[end] {
		return bs(nums, middle+1, end, target)
	}
	return bs(nums, start, middle-1, target)
}

func search(nums []int, target int) int {
	return bs(nums, 0, len(nums)-1, target)
}

func main() {

	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 1
	fmt.Print("Output ", search(nums, target))

}
