package main

import (
	"fmt"
)

//滑动窗口 思路就是维持一个不重复的窗口一直移动，这里的lookup 和charset就是对应的窗口，然后内部有一个
//窗口的扫描器，做窗口内部的子扫描

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	lookup := make(map[byte]bool)
	maxStr := 0
	left := 0
	for i := 0; i < len(s); i++ {
		for lookup[s[i]] {
			delete(lookup, s[left])
			left++
		}
		if maxStr < i-left+1 {
			maxStr = i - left + 1
		}
		lookup[s[i]] = true
	}
	return maxStr
}

func main() {
	test := "abc"
	fmt.Println(lengthOfLongestSubstring(test)) // 输出: 3
	fmt.Println(LongestSubstring(test))
}

func LongestSubstring(s string) string {
	charSet := make(map[byte]bool)
	left := 0
	maxStr := ""

	for i := 0; i < len(s); i++ {
		for charSet[s[i]] {
			delete(charSet, s[left])
			left++
		}

		if i-left+1 > len(maxStr) {
			maxStr = s[left : i+1]
		}
		charSet[s[i]] = true

	}

	return maxStr
}
