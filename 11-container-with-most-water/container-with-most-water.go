package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1}))
}

func maxArea(height []int) int {
	l := len(height)

	left, right := 0, l-1
	ans := 0

	for left < right {
		// 计算面积
		sq := min(height[left], height[right]) * (right - left)
		if sq > ans {
			ans = sq
		}
		// 如果左板大于右板，则右板收窄，否则相反
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
