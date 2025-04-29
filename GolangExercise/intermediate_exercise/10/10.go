//Question£º Write a program that accepts sequence of lines as input and prints the lines after making all characters in the sentence capitalized. Suppose the following input is supplied to the program: Hello world Practice makes perfect Then, the output should be: HELLO WORLD PRACTICE MAKES PERFECT

package main

import (
	"fmt"
	"strings"
)

func capitalizedWord(char string) string {
	return strings.ToUpper(char)
}

func main() {
	var a string
	for _, char := range "Hello world Practice makes perfect" {
		a += capitalizedWord(string(char))
	}
	fmt.Print(a)
}
