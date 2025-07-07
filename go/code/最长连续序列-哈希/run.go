package main

import "fmt"

func longestConsecutive(nums []int) (ans int) {
	hashMap := make(map[int]bool)
	for _, item := range nums {
		hashMap[item] = true
	}

	for k := range hashMap {
		if hashMap[k-1] {
			continue
		}

		y := k + 1

		for hashMap[y] {
			y++
		}

		ans = max(ans, y-k)
	}

	return ans
}

func main() {
	var src = []int{1, 2, 3, 55, 3}
	fmt.Println(longestConsecutive(src))
}
