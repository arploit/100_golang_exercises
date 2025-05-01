package main

/*
	- Search First Tree
	- Backtracking
	- DFS
	- Decision making

*/

import (
	"fmt"
	"time"
)

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var dfs func(index int, curr []int, total int)
	dfs = func(index int, curr []int, total int) {

		if total == target {
			combination := append([]int{}, curr...)
			result = append(result, combination)
			return
		}
		if total > target || index >= len(candidates) {
			return
		}

		dfs(index, append(curr, candidates[index]), total+candidates[index])
		dfs(index+1, curr, total)

	}

	dfs(0, []int{}, 0)
	return result
}

func main() {
	start := time.Now()
	candidates := []int{2, 3, 5}
	target := 8

	fmt.Print(combinationSum(candidates, target))
	fmt.Println(time.Since(start))
}
