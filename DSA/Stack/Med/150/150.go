// Q : https://leetcode.com/problems/evaluate-reverse-polish-notation/description/

package main

import (
	"fmt"
	"strconv"
)

func performCalculation(stack []int, operation string) []int {
	if len(stack) >= 2 {

		lastEle := stack[len(stack)-1]
		sLastEle := stack[len(stack)-2]
		result := 0

		switch operation {
		case "+":
			result = lastEle + sLastEle
		case "-":
			result = sLastEle - lastEle
		case "*":
			result = lastEle * sLastEle

		case "/":
			result = sLastEle / lastEle
		default:
			return stack
		}

		stack = stack[:len(stack)-2]
		stack = append(stack, result)

	}
	return stack

}

func evalRPN(tokens []string) int {
	if len(tokens) >= 3 {
		stack := []int{}
		for i := range tokens {
			element := tokens[i]
			num, err := strconv.ParseInt(element, 10, 32)
			if err != nil {
				stack = performCalculation(stack, element)
			} else {
				stack = append(stack, int(num))
			}

		}
		return stack[len(stack)-1]
	}
	num, _ := strconv.Atoi(tokens[len(tokens)-1])
	return num
}

func main() {
	tokens := []string{"3", "-4", "+"}

	fmt.Println("Output", evalRPN(tokens))
}
