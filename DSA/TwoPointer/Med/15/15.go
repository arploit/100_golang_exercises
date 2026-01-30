package main

import (
	"fmt"
	"sort"
)

// Q : https://leetcode.com/problems/3sum/description/

func threeSum(nums []int) [][]int {

	var result [][]int
	if len(nums) < 3 {
		return result
	}
	seen := make(map[string]bool)
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := (len(nums) - 1)
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				key := fmt.Sprintf("%d,%d,%d", nums[i], nums[left], nums[right])
				if !seen[key] {
					seen[key] = true
					result = append(result, []int{nums[i], nums[left], nums[right]})
				}
				left++
				right--
			} else if sum < 0 {
				left += 1
			} else if sum > 0 {
				right -= 1
			}
		}
	}

	return result

}

func main() {
	arr := []int{-1, 0, 1, 2, -1, -4}

	fmt.Printf("%v", threeSum(arr))

}
