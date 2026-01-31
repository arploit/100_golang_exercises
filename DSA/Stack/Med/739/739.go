package main

import "fmt"

// Q : https://leetcode.com/problems/daily-temperatures/

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)

	stack := []int{}
	for i := 0; i <= len(temperatures)-1; i++ {

		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			prevDay := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[prevDay] = i - prevDay
		}

		stack = append(stack, i)
	}

	return result

}

func main() {
	temperature := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println("Output", dailyTemperatures(temperature))

}

//Expected O/p
// [1,1,4,2,1,1,0,0]
