//Question: Write a program that accepts a comma separated sequence of words as
// input and prints the words in a comma-separated sequence after sorting them alphabetically.\
//  Suppose the following input is supplied to the program: without,hello,bag,world Then, the output should be:
//  bag,hello,without,world

// Hints: In case of input data being supplied to the question, it should be assumed to be a console input.

package main

import (
	"fmt"
	"strings"
)

func swap(left []string, right []string) []string {
	i, j := 0, 0
	leftLength := len(left)
	rightLength := len(right)
	var newArray []string
	if i < leftLength && j < rightLength {
		if left[i] > right[j] {
			newArray = append(newArray, right[j])
			j++
		} else {
			newArray = append(newArray, left[i])
			i++
		}
	}

	if i < leftLength {
		newArray = append(newArray, left[i:]...)
	}

	if j < rightLength {
		newArray = append(newArray, right[j:]...)
	}

	return newArray

}

func mergeSort(wordArray []string) []string {
	if len(wordArray) <= 1 {
		return wordArray
	}

	left := mergeSort(wordArray[:len(wordArray)/2])
	right := mergeSort(wordArray[len(wordArray)/2:])

	newArray := swap(left, right)

	return newArray
}

func sortSentence(sentence string) {
	v := strings.Split(sentence, ",")
	z := mergeSort(v)
	for idx, word := range z {
		if idx == len(z)-1 {

			fmt.Printf("%s", word)
		} else {

			fmt.Printf("%s,", word)
		}
	}

}

func main() {
	v := "without,hello,bag,world,Then"
	sortSentence(v)

}
