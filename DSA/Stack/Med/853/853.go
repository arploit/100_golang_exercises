// Q https://leetcode.com/problems/car-fleet/description/

package main

import (
	"fmt"
	"sort"
)

func carFleet(target int, position []int, speed []int) int {

	posSped := make([][]int, len(position))

	stack := []float64{}
	for i := range position {
		posSped[i] = []int{position[i], speed[i]}
	}
	sort.Slice(posSped, func(i, j int) bool {
		return posSped[i][0] > posSped[j][0]
	})

	for i := range posSped {
		time := float64(target-posSped[i][0]) / float64(posSped[i][1])
		stack = append(stack, (time))
		if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack)
}

func main() {
	target := 10
	position := []int{6, 8}
	speed := []int{3, 2}

	fmt.Println("Output", carFleet(target, position, speed))

}
