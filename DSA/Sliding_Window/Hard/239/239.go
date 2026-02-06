// Q: https://leetcode.com/problems/sliding-window-maximum/
package main

import "fmt"

/*
Explaination : Need to use dequeue
we check every element one by one and push them into the dequeue as they are encountered and remove them from the front if window has moved from their position.

We need to pop all the element from the back untill current index value is bigger then the value present inside the deque. That could mean either its empty or I found the largest number than current value at index i


for a result only starts pushing untill i has reached to the edge of window and we found the largest element present at the front dequeue


*/

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	result := []int{}
	d := []int{}

	for i := 0; i < len(nums); i++ {
		if len(d) > 0 && d[0] <= i-k {
			d = d[1:]
		}

		for len(d) > 0 && nums[d[len(d)-1]] <= nums[i] {
			d = d[:len(d)-1]
		}

		d = append(d, i)

		if i >= k-1 {
			result = append(result, nums[d[0]])
		}
	}

	return result
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println("Output", maxSlidingWindow(nums, k))
}
