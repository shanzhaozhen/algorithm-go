package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 2}))
}

func maxArea(height []int) int {
	left, right, ans := 0, len(height)-1, 0
	for left < right {
		sq := min(height[left], height[right]) * (right - left)
		if sq > ans {
			ans = sq
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
