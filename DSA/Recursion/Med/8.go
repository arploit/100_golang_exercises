package main

// Q : https://leetcode.com/problems/string-to-integer-atoi/description/

import (
	"fmt"
	"math"
)

func atoi(s string) int {
	result := 0
	sign := 1
	started := false

	for i := 0; i < len(s); i++ {
		c := s[i]

		if !started {
			if c == ' ' {
				continue
			} else if c == '+' {
				sign = 1
				started = true
			} else if c == '-' {
				sign = -1
				started = true
			} else if c >= '0' && c <= '9' {
				started = true
				digit := int(c - '0')
				result = digit
			} else {
				return 0
			}
		} else {
			if c >= '0' && c <= '9' {
				digit := int(c - '0')

				if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > 7) {
					if sign == 1 {
						return math.MaxInt32
					} else {
						return math.MinInt32
					}
				}

				result = result*10 + digit
			} else {
				break
			}
		}
	}

	return result * sign
}

func main() {
	s := "      -042"

	fmt.Println("Output :", atoi(s))
}
