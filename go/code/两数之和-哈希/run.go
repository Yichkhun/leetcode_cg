package main

func twoSum(src []int, target int) []int {
	var result []int
	hashMap := make(map[int64]int)
	for index, item := range src {
		aim := target - item
		if val, find := hashMap[int64(aim)]; find {
			result = append(result, index, val)
		}
		hashMap[int64(item)] = index
	}

	return result
}

func main() {

}
