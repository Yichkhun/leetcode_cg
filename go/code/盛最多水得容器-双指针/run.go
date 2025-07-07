package main

import "math"

func main() {

}

func maxArea(height []int) int64 {
	if len(height) == 0 {
		return 0
	}
	left := 0
	right := len(height) - 1
	rea := int64(0)

	for left < right {
		_rea := int64(math.Min(float64(height[left]), float64(height[right]))) * int64(right-left)
		rea = int64(math.Max(float64(_rea), float64(rea)))

		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return rea
}
