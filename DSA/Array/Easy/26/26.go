//https://leetcode.com/problems/remove-duplicates-from-sorted-array

package main

import "fmt"

func removeDuplicates(nums []int) int {
	hash := make(map[int]int)
	idx := 0
	if len(nums) <= 1 {
		return len(nums)
	}
	for _, num := range nums {
		if hash[num] == 0 {
			hash[num] = 1
			nums[idx] = num
			idx++
		}
	}

	for i := idx; i < len(nums); i++ {
		nums[i] = 99
	}

	return idx

}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums), nums)
}
