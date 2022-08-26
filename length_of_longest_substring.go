package main

import (
	"fmt"
)

// 不重复的最长子串长度
func lengthOfLongestSubstring(s string) int {
	freq := make(map[byte]int, 256)
	l, r := 0, -1
	res := 0
	for l < len(s) {
		if r < len(s)-1 && freq[s[r+1]] == 0 {
			r++
			freq[s[r]]++
		} else {
			freq[s[l]]--
			l++
		}
		if r-l+1 > res {
			res = r - l + 1
		}
	}
	return res
}

func main() {
	a := []string{
		"abcabcbb", // 3
		"bbbbb",    // 1
		"pwwkew",   // 3
		"",         // 0
	}
	for _, str := range a {
		fmt.Printf("%q: %d\n", str, lengthOfLongestSubstring(str))
	}
}
