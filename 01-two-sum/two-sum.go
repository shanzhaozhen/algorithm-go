package main

import (
	"fmt"
)

func main() {
	sum := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(sum)
}

/*
*
冒泡计算
*/
func twoSum(nums []int, target int) []int {
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

/*
*
哈希表
*/
func twoSum1(nums []int, target int) []int {
	hashTable := map[int]int{}

	for i, num := range nums {
		// 查看当前值与目标值之差的值在不在哈希表中，在则是之前已经遍历过了，直接拿之前的索引
		if p, ok := hashTable[target-num]; ok {
			return []int{p, i}
		}
		// 不在索引则记录当前值的索引
		hashTable[num] = i
	}

	return nil
}
