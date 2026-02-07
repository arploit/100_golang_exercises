// Q:// https://leetcode.com/problems/search-a-2d-matrix/description/
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

func searchInside(mat [][]int, start int, end int, target int) bool {

	if start > end {
		return false
	}
	mid := ((end - start) / 2) + start
	result := -1

	if mat[mid][0] <= target && mat[mid][len(mat[mid])-1] >= target {
		result = bs(mat[mid], 0, len(mat[mid])-1, target)
		if result != -1 {
			return true
		}
	}

	if target > mat[mid][len(mat[mid])-1] {
		return searchInside(mat, mid+1, end, target)
	}

	if target < mat[mid][0] {
		return searchInside(mat, start, mid-1, target)
	}

	return false

}

func searchMatrix(matrix [][]int, target int) bool {
	return searchInside(matrix, 0, len(matrix)-1, target)
}

func main() {

	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 3

	fmt.Println("Output", searchMatrix(matrix, target))
}
