package main

import "fmt"

func main() {
	fmt.Println(threeSum([]int{1, -1, 0}))
}

func threeSum(nums []int) [][]int {
	l := len(nums)

	var ans [][]int

	for i := 0; i < l; i++ {
		// 需要和上一次枚举的数不相同
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					ans = append(ans, []int{nums[i], nums[j], nums[k]})
				}
			}
		}

	}
	return ans
}
