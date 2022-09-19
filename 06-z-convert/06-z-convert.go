package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}

// 利用二维矩阵模拟
func convert(s string, numRows int) string {
	n, r := len(s), numRows
	if r == 1 || r >= n {
		return s
	}
	// 每个周期的字符数（横纵各是行数，减去转角位重复和首尾相接的一位）
	t := r*2 - 2
	// 矩阵最大列数
	c := (n + t - 1) / t * (r - 1)
	mat := make([][]byte, r)
	for i := range mat {
		mat[i] = make([]byte, c)
	}
	x, y := 0, 0
	for i, ch := range s {
		mat[x][y] = byte(ch)
		// 当前周期小于行数时
		if i%t < r-1 {
			// 向下移动
			x++
		} else {
			// 向右上移动
			x--
			y++
		}
	}
	ans := make([]byte, 0, n)
	for _, row := range mat {
		for _, ch := range row {
			if ch > 0 {
				ans = append(ans, ch)
			}
		}
	}
	return string(ans)
}

// 压缩矩阵空间
func convert2(s string, numRows int) string {
	r := numRows
	if r == 1 || r >= len(s) {
		return s
	}
	mat := make([][]byte, r)
	t := r*2 - 2
	x := 0
	for i, ch := range s {
		// y 轴活动是单向的尾插入字符串即可
		mat[x] = append(mat[x], byte(ch))
		if i%t < r-1 {
			x++
		} else {
			x--
		}
	}
	return string(bytes.Join(mat, nil))
}

// 直接构造
func convert3(s string, numRows int) string {
	n, r := len(s), numRows
	if r == 1 || r >= n {
		return s
	}
	t := r*2 - 2
	ans := make([]byte, 0, n)
	for i := 0; i < r; i++ { // 枚举矩阵的行
		for j := 0; j+i < n; j += t { // 枚举每个周期的起始下标
			ans = append(ans, s[j+i]) // 当前周期的第一个字符
			if 0 < i && i < r-1 && j+t-i < n {
				ans = append(ans, s[j+t-i]) // 当前周期的第二个字符
			}
		}
	}
	return string(ans)
}
