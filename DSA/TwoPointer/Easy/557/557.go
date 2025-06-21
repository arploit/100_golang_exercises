//https://leetcode.com/problems/reverse-words-in-a-string-iii/description/?envType=problem-list-v2&envId=two-pointers

package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	splittedString := strings.Fields(s)
	var builder strings.Builder

	for _, word := range splittedString {
		for i := len(word) - 1; i >= 0; i-- {
			builder.WriteByte(word[i])
		}
		builder.WriteByte(' ')
	}

	return strings.TrimSpace(builder.String())
}

func main() {

	word := "Let's take LeetCode contest"
	newWord := reverseWords(word)
	fmt.Println(newWord)

}
