// Q: // https://leetcode.com/problems/koko-eating-bananas/
package main

import (
	"fmt"
	"slices"
)

func searchOptimiseAns(piles []int, start int, end int, h int) int {
	if start > end {
		return start
	}

	middle := ((end - start) / 2) + start
	sum := 0

	for i := range piles {
		if middle >= piles[i] {
			sum += 1
		} else {
			if piles[i]%middle != 0 {
				sum += piles[i]/middle + 1
			} else {
				sum += piles[i] / middle
			}
		}
	}

	if sum <= h {
		return searchOptimiseAns(piles, start, middle-1, h)
	} else {
		return searchOptimiseAns(piles, middle+1, end, h)
	}

}

func minEatingSpeed(piles []int, h int) int {

	max := slices.Max(piles)

	return searchOptimiseAns(piles, 1, max, h)

}

func main() {
	piles := []int{3, 6, 7, 11}
	h := 8

	fmt.Println("Output", minEatingSpeed(piles, h))
}
