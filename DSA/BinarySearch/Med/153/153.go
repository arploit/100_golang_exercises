// Q: // https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/

package main

import (
	"fmt"
)

func bs(nums []int, start int, end int) int {
	if start >= end {
		return nums[start]
	}

	middle := ((end - start) / 2) + start

	if nums[middle] > nums[end] {
		return bs(nums, middle+1, end)
	} else {
		return bs(nums, start, middle)
	}
}

func findMin(nums []int) int {
	return bs(nums, 0, len(nums)-1)
}

func main() {

	nums := []int{2, 1} //
	fmt.Print("Output", findMin(nums))

}
