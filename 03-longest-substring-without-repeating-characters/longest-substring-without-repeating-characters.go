package main

func main() {
}

func lengthOfLongestSubstring(s string) int {
	hashTable := map[byte]int{}

	ans, left := 0, 0
	for i := 0; i < len(s); i++ {
		// 当前字符如果在哈希表已经出现，则与之前记录最大字段的位置做对比，记录当前子串的开始位置
		if _, ok := hashTable[s[i]]; ok {
			left = max(left, hashTable[s[i]]+1)
		}
		// 记录当前字符最后一次出现该字段的位置
		hashTable[s[i]] = i

		// 记录最长子串长度
		ans = max(ans, i-left+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
