package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{1, -1, 0}))
}

func threeSum(nums []int) [][]int {
	l := len(nums)
	// 先排序
	sort.Ints(nums)
	ans := make([][]int, 0)

	for i := 0; i < l; i++ {
		// 需要和上一次枚举的数不相同
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		k := l - 1
		for j := i + 1; j < l; j++ {
			// 需要和上一次枚举的数不相同
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			// 需要保证 b 的指针在 c 的指针的左侧
			// 如果 a + b + c > 0 则需要往左移动
			for j < k && nums[i]+nums[j]+nums[k] > 0 {
				k--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if j == k {
				break
			}

			if nums[i]+nums[j]+nums[k] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
			}
		}

	}
	return ans
}

func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	l := len(nums)

	ans := make([][]int, 0)

	for i := 0; i < l; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j, k := i+1, l-1; j < l && k < l; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			for ; j < k && nums[i]+nums[j]+nums[k] > 0; k-- {
			}

			if j == k {
				break
			}

			if nums[i]+nums[j]+nums[k] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
			}
		}

	}

	return ans

}
