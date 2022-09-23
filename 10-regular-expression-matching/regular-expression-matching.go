package main

import "fmt"

func main() {
	fmt.Println(isMatch("aa", ".x"))
}

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)

	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	// dp[i][j] 表示 s 的前 i 个是否能被 p 的前 j 个匹配
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 如果 p 当前位置的前一位是 *,则可以匹配当前 s 的位置
			if p[j-1] == '*' {
				// 查看 p 中 * 的前一位能不能匹配上
				dp[i][j] = dp[i][j] || dp[i][j-2]
				//
				if matches(i, j-1) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else if matches(i, j) {
				dp[i][j] = dp[i][j] || dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}
