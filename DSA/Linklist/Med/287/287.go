package main

import "fmt"

//287 : https://leetcode.com/problems/find-the-duplicate-number/description/
func findDuplicate(nums []int) int {
	i := 0
	j := 0

	for true {
		i = nums[i]
		j = nums[nums[j]]
		if i == j {
			break
		}
	}

	j = 0
	for nums[i] != nums[j] {
		i = nums[i]
		j = nums[j]
	}

	return nums[i]

}

func main() {
	array := [5]int{1, 3, 4, 2, 2}
	fmt.Println("Repeated Number : ", findDuplicate(array[:]))

}
