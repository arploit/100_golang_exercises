package main

import "fmt"

// Q https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/

func maxProfit(prices []int) int {
	if len(prices) >= 2 {
		maxProfit := 0
		var minPrice int = prices[0]
		for i := range prices {

			if minPrice >= prices[i] {
				minPrice = prices[i]
			}

			if maxProfit <= (prices[i] - minPrice) {
				maxProfit = prices[i] - minPrice

			}
		}

		return maxProfit
	}
	return 0
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Print("Output ", maxProfit(prices))
}
