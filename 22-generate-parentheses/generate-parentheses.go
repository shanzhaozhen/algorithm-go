package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	var res []string

	if n == 0 {
		return res
	}

	dfs("", n, n, &res)
	return res
}

func dfs(curStr string, leftNum int, rightNum int, res *[]string) {
	// 递归终止直接加入结果集，无需回溯
	if leftNum == 0 && rightNum == 0 {
		*res = append(*res, curStr)
		return
	}

	// 剪枝
	if leftNum > rightNum {
		return
	}

	if leftNum > 0 {
		dfs(curStr+"(", leftNum-1, rightNum, res)
	}

	if rightNum > 0 {
		dfs(curStr+")", leftNum, rightNum-1, res)
	}
}
