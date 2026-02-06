// Q: https://leetcode.com/problems/longest-repeating-character-replacement/description/

package main

import "fmt"

func characterReplacement(s string, k int) int {
	freq := [26]int{}
	left, maxlen, maxFreq := 0, 0, 0
	for right := 0; right < len(s); right++ {
		index := s[right] - 'A'
		freq[index]++

		if freq[index] > maxFreq {
			maxFreq = freq[index]
		}

		if (right-left+1)-maxFreq > k {
			freq[s[left]-'A']--
			left++
		}

		if right-left+1 > maxlen {
			maxlen = right - left + 1
		}

	}

	return maxlen
}

func main() {
	s := "ABAB"
	k := 2

	fmt.Println("Output", characterReplacement(s, k))
}
