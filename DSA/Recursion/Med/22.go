package main

// Q : https://leetcode.com/problems/generate-parentheses/description/

import "fmt"

type Stack struct {
	elements []interface{}
}

func generateParenthesis(n int) []string {
	stack := []string{}
	result := []string{}

	var backtrack func(openN int, closeN int)
	backtrack = func(openN int, closeN int) {

		if openN == n && closeN == n {
			// join stack into a string
			curr := ""
			for _, ch := range stack {
				curr += ch
			}
			result = append(result, curr)
			return
		}

		if openN < n {
			stack = append(stack, "(")
			backtrack(openN+1, closeN)
			stack = stack[:len(stack)-1]
		}

		if closeN < openN {
			stack = append(stack, ")")
			backtrack(openN, closeN+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)
	return result
}

func main() {

	fmt.Printf("%v\n", generateParenthesis(3))
}
