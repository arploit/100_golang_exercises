// Q : https://leetcode.com/problems/minimum-window-substring/description/

package main

import "fmt"

func minWindow(s string, t string) string {

	leftcur, rightCur, left, minWin := 0, 0, 0, len(s)
	h1, h2 := [128]int{}, [128]int{}

	for i := 0; i < len(t); i++ {
		h1[t[i]]++
	}

	for right := 0; right < len(s); right++ {
		h2[s[right]]++

		if right-left+1 < len(s) && right-left+1 >= len(t) {
			for left < right {
				if h1[s[left]] == 0 {
					left++
				} else if h2[s[left]] > h1[t[left]] && h1[s[left]] > 0 {
					h2[s[left]]--
					left++
				} else {
					break
				}
			}

			if right-left+1 <= minWin {
				leftcur = left
				rightCur = right
				minWin = right - left + 1
			}
		}

	}

	if minWin <= len(s) && minWin >= len(t) {
		return s[leftcur:rightCur]
	}

	return ""
}

func main() {

	s := "ADOBECODEBANC"
	t := "ABC"

	fmt.Println("Output", minWindow(s, t))

}
