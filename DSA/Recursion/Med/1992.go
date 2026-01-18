package main

import (
	"fmt"
)

// Q : https://leetcode.com/problems/count-good-numbers/description/

func myPow(x int, n int64) int {

	mod := (1000000007)

	if n == 0 {
		return 1
	}

	if n == 1 {
		return (x)
	}

	half := myPow(x, n/2) % mod

	if n%2 == 0 {
		return (half * half) % mod
	}

	return (half * half * (x)) % mod

}

func countGoodNumbers(n int64) int {
	mod := (1000000007)
	if n%2 == 0 {
		return (myPow(5, n/2) * myPow(4, n/2)) % mod

	}

	return (myPow(5, ((n+1)/2)) * myPow(4, (n/2))) % mod

}

func main() {

	fmt.Println(countGoodNumbers(1))
	fmt.Println(countGoodNumbers(2))
	fmt.Println(countGoodNumbers(5))
	fmt.Println(countGoodNumbers(50))

	// fmt.Println(countGoodNumbers(1))

}
