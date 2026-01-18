package main

import "fmt"

// Q : https://leetcode.com/problems/powx-n/description/

func myPow(x float64, n int) float64 {

	if x == 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	if n < 0 {
		return (1 / myPow(x, -n))
	}

	half := myPow(x, n/2)

	if n%2 == 0 {
		return half * half
	}

	return half * half * x

}

func main() {

	x := 2.00000
	n := -200000000

	fmt.Println(myPow(x, n))
}
