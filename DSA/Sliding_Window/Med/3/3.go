package main

// Q https://leetcode.com/problems/longest-substring-without-repeating-characters/
import "fmt"

func lengthOfLongestSubstring(s string) int {

	hashmap := make(map[rune]int)
	left, maxLength := 0, 0

	for idx, ch := range s {
		if prev, ok := hashmap[ch]; ok && prev >= left {
			left = prev + 1
		}

		hashmap[ch] = idx

		if curr := idx - left + 1; curr > maxLength {
			maxLength = curr
		}
	}

	return maxLength
}

func main() {
	s := "abcabcbb"

	fmt.Println("Output", lengthOfLongestSubstring(s))
}
