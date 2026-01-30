package main

import "fmt"

// Q : https://leetcode.com/problems/min-stack/

type MinStack struct {
	stack []int
	min   []int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		min:   []int{},
	}
}

func (this *MinStack) Push(val int) {
	stack := this.stack
	minA := this.min
	stack = append(stack, val)
	this.stack = stack
	if len(minA) > 0 {
		currentMinimum := minA[len(minA)-1]
		minimum := min(currentMinimum, val)
		minA = append(minA, minimum)
	} else {
		minA = append(minA, val)
	}
	this.min = minA
}

func (this *MinStack) Pop() {
	if len(this.stack) > 0 {
		this.stack = this.stack[:len(this.stack)-1]
		this.min = this.min[:len(this.min)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) > 0 {
		return this.stack[len(this.stack)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if len(this.min) > 0 {
		return this.min[len(this.min)-1]
	}
	return 0
}

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	fmt.Println(minStack.GetMin()) // expected: -3

	minStack.Pop()

	fmt.Println(minStack.Top())    // expected: 0
	fmt.Println(minStack.GetMin()) // expected: -2

}
