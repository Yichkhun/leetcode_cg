package main

import "fmt"

func findSumIndex(src []int64, target int64) []int64 {
	var result []int64
	hashMap := make(map[int64]int64)
	for index, item := range src {
		aim := target - item
		if val, find := hashMap[aim]; find {
			result = append(result, int64(index), val)
		}
		hashMap[item] = int64(index)
	}

	return result
}

func main() {
	var nums = []int64{2, 7, 11, 15}
	target := 9
	fmt.Println(findSumIndex(nums, int64(target)))
}
