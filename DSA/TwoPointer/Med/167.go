// Q : https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	/*
		Complete Solution Using O(n2)

		// arrayLength := len(numbers)
		// result := []int{}

		// // for i < arrayLength {
		// // 	fmt.Println(numbers[i], numbers[i] <= target)
		// // 	if numbers[i] <= target {
		// // 		break
		// // 	}
		// // 	i += 1
		// // }

		// a1 := 0
		// a2 := 0

		// for a1 < arrayLength {
		// 	a2 = a1 + 1
		// 	for a2 < arrayLength {
		// 		fmt.Println(numbers[a1], numbers[a2], numbers[a1]+numbers[a2])
		// 		if numbers[a1]+numbers[a2] == target {
		// 			result = append(result, a1+1, a2+1)
		// 			return result
		// 		}
		// 		a2 += 1
		// 	}
		// 	a1 += 1
		// }

		// return result
	*/

	hashmap := make(map[int]int)

	for i, num := range numbers {
		if idx, ok := hashmap[target-num]; ok {
			return []int{idx, i + 1}
		}
		hashmap[num] = i + 1
	}
	return []int{}

}

func main() {
	numbers := []int{2, 3, 4}
	target := 6
	fmt.Printf("%v", twoSum(numbers, target))
}
