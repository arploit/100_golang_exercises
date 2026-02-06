// Q https://leetcode.com/problems/permutation-in-string/description/
package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	h1 := [26]int{0}
	h2 := [26]int{0}

	left := 0
	for i := 0; i < len(s1); i++ {
		h1[s1[i]-'a']++
	}

	for right := 0; right < len(s2); right++ {

		h2[s2[right]-'a']++

		if right-left+1 > len(s1) {
			h2[s2[left]-'a']--
			left++
		}

		if right-left+1 == len(s1) && h1 == h2 {
			return true
		}

	}

	return false
}

func main() {
	s1 := "ab"
	s2 := "eidbaooo"

	fmt.Println("oUTPUT", checkInclusion(s1, s2))

}
