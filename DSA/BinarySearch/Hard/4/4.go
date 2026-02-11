// Q:https://leetcode.com/problems/median-of-two-sorted-arrays/

package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1 := len(nums1)
	n2 := len(nums2)

	if n1 > n2 {
		return findMedianSortedArrays(nums2, nums1)
	}
	low := 0
	high := n1
	total := n1 + n2
	mid := (n1 + n2 + 1) / 2
	for true {
		mid1 := (low + high) / 2
		mid2 := mid - mid1
		l1 := math.Inf(-1)
		l2 := math.Inf(-1)
		r1 := math.Inf(1)
		r2 := math.Inf(1)
		if mid1 < n1 {
			r1 = float64(nums1[mid1])
		}
		if mid2 < n2 {
			r2 = float64(nums2[mid2])
		}

		if mid1-1 >= 0 {
			l1 = float64(nums1[mid1-1])
		}

		if mid2-1 >= 0 {
			l2 = float64(nums2[mid2-1])
		}

		if l1 <= r2 && l2 <= r1 {
			if total%2 == 0 {
				return (math.Max(l1, l2) + math.Min(r1, r2)) / 2
			} else {
				return math.Max(l2, l1)
			}
		} else if l1 > r2 {
			high = mid1 - 1
		} else {
			low = mid1 + 1
		}
	}
	return 0
}

func main() {

	num1 := []int{1, 3} //7, 12, 14, 15
	num2 := []int{2}    //

	fmt.Println("Output", findMedianSortedArrays(num1, num2))

}
