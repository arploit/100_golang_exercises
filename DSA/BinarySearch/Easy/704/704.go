// Q: // https://leetcode.com/problems/binary-search/description/

package main

import "fmt"

func bs(nums []int, start int, end int, target int) int {

	if start > end {
		return -1
	}

	middle := ((end - start) / 2) + start

	if nums[middle] == target {
		return middle
	}

	if nums[middle] >= target {
		return bs(nums, start, middle-1, target)
	}

	if nums[middle] < target {
		return bs(nums, middle+1, end, target)

	}

	return -1
}

func search(nums []int, target int) int {
	end := len(nums) - 1
	return bs(nums, 0, end, target)
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9

	fmt.Print("Output ", search(nums, target))

}
