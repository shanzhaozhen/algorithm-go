package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome2("aacabdkacaa"))
}

/*
*
动态规划
*/
func longestPalindrome(s string) string {
	L := len(s)
	if L < 2 {
		return s
	}

	maxLen := 1
	begin := 0

	// dp[i][j] 表示 s[i..j] 是否是回文串
	var dp = make([][]bool, L)
	// 初始化：所有长度为 1 的子串都是回文串
	for i := range dp {
		dp[i] = make([]bool, L)
		dp[i][i] = true
	}

	/**
	枚举子串长度(l)
	左边界（i）
	右边界（j）
	*/
	for l := 1; l <= L; l++ {
		// 枚举左边界，左边界的上限设置可以宽松一些
		for i := 0; i < L; i++ {
			// 由 i 和 j 可以确定右边界，即 j - i + 1 = l 得
			j := i + l - 1
			// 如果右边界越界，就可以退出当前循环
			if j >= L {
				break
			}

			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					// 两边收窄一位，字串是回文串才成立
					dp[i][j] = dp[i+1][j-1]
				}
			}

			// 只要 dp[i][j] == true 成立，就表示子串 s[i..j] 是回文，此时记录回文长度和起始位置
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}

	return s[begin : begin+maxLen]
}

func longestPalindrome2(s string) string {
	L := len(s)
	if L <= 1 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < L; i++ {
		// 分别拿到奇数偶数的回文子串长度
		// 奇数中心的扩散长度
		left1, right1 := expandAroundCenter(s, i, i)
		// 偶数中心的扩散长度
		left2, right2 := expandAroundCenter(s, i, i+1)
		// 对比最大的长度
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
