package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	record := make(map[string][]string)
	for _, item := range strs {
		temp := item
		st := []byte(temp)
		sort.Slice(st, func(i, j int) bool {
			return st[i] > st[j]
		})
		record[string(st)] = append(record[string(st)], item)
	}

	res := make([][]string, len(record))
	index := 0
	for _, v := range record {
		res[index] = append(res[index], v...)
		index++
	}

	return res
}

func main() {
	var strs = []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))

}
