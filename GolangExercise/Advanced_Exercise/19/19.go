// https://github.com/zhiwehu/Python-programming-exercises/blob/master/100%2B%20Python%20challenging%20programming%20exercises%20for%20Python%203.md#question-19

package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

func SortStudent (students []Student){
	n := len(students)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if  students[j].Name > students[j+1].Name ||
			(students[j].Name == students[j+1].Name && students[j].Age > students[j+1].Age) ||
			(students[j].Name == students[j+1].Name && students[j].Age == students[j+1].Age && students[j].Score > students[j+1].Score) {
			 
				// Swap students[j] and students[j+1]
				students[j], students[j+1] = students[j+1], students[j]
			}
			
		}
	}
	
}

func main() {
	var students = []Student{
		{"John", 19, 90},
		{"Alice", 20, 85},
		
		{"Bob", 18, 78},
		{"Diana", 21, 92},
		{"Alice", 21, 85},
		{"Alice", 20, 83},

	}
	fmt.Println("Before Sort:",students)
	SortStudent(students)
	fmt.Println("After Sort:",students)
}