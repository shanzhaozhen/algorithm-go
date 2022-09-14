package main

import (
	"fmt"
	"math"
)

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{1, 2}
	arrays := findMedianSortedArrays(nums1, nums2)
	fmt.Println(arrays)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 计算总元素个数
	totalLength := len(nums1) + len(nums2)

	// 奇偶情况区分处理，用二分法找第K小的数
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		// 如果 index1 已到 数组1 最大长度，则中位数在 数组2 中，仅需在 数组2 当前游标开始算的第k个数即可
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		// 同上
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		// 中位数在剩下所有字段的第一个，所以取最小值即可
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		/**
		1. 假如第一数组的中位数比第二数组的中位数小
			则下一次循环从第一个数组的中位数开始找，由于范围缩小，减少了中位数前面的数：newIndex1 - index1 + 1，
			所以找第 k - (newIndex1 - index1 + 1) 小的数即可
			更新第一数组的开始位置 index1 = newIndex1 + 1
		2. 与1情况一样
		*/
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1
		} else {
			k -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	m, n := len(nums1), len(nums2)
	left, right := 0, m
	median1, median2 := 0, 0
	for left <= right {
		i := (left + right) / 2
		j := (m+n+1)/2 - i
		numsIm1 := math.MinInt32
		if i != 0 {
			numsIm1 = nums1[i-1]
		}
		numsI := math.MaxInt32
		if i != m {
			numsI = nums1[i]
		}
		numsJm1 := math.MinInt32
		if j != 0 {
			numsJm1 = nums2[j-1]
		}
		numsJ := math.MaxInt32
		if j != n {
			numsJ = nums2[j]
		}
		if numsIm1 <= numsJ {
			median1 = max(numsIm1, numsJm1)
			median2 = min(numsI, numsJ)
			left = i + 1
		} else {
			right = i - 1
		}
	}
	if (m+n)%2 == 0 {
		return float64(median1+median2) / 2.0
	}
	return float64(median1)
}
