package main

import "fmt"

func main() {
	var src = []int{1, 0, 4, 7, 0, 10, 11}
	moveZeroes(src)
	fmt.Println(src)
}

func moveZeroes(nums []int) {
	left, right := 0, 0
	length := len(nums)

	for i := 0; i < length; i++ {
		if nums[right] != 0 {
			_tmp := nums[right]
			nums[right] = nums[left]
			nums[left] = _tmp
			left++
		}
		right++
	}
}
