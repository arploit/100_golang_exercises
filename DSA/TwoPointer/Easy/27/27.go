package main

import "fmt"

func removeElement(nums []int, val int) int {
	left, k := 0, 0
	right := len(nums) - 1

	for left <= right {
		if nums[left] == val && nums[right] != val {
			nums[left], nums[right] = nums[right], nums[left]
			k++
			left++
			right--
		}

		if nums[left] != val {
			k++
			left++
		}

		if nums[right] == val {
			right--
		}

		if left == right && nums[left] != val {
			k++
			left++
			right--
		}

	}
	return k
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	newLength := removeElement(nums, val)
	fmt.Print("New length of the array is: ", newLength)
	fmt.Println(nums)
}
