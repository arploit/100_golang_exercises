// 56. Merge Intervals
// https://leetcode.com/problems/merge-intervals/description/?envType=problem-list-v2&envId=array

/*
-Sort the array

- Visualise the number line

- find start of current array with last added value [1] and find which one is bigger and update that value with last added value

if not append the array in output
*/
package main

import "fmt"

func merge(intervals [][]int) [][]int {

}

func main() {

	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}

	fmt.Println(merge(intervals))
}
