package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("42"))
}

func myAtoi(s string) int {
	i, abs := 0, 0
	l := len(s)
	sign := 1

	// 去前空格
	for i = 0; i < l && s[i] == ' '; i++ {
	}

	if i >= l {
		return 0
	}

	// 检查正负号
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}

	for ; i < l && s[i] >= '0' && s[i] <= '9'; i++ {
		abs = abs*10 + int(s[i]-'0')
		if sign*abs < math.MinInt32 { //整数超过 32 位有符号整数范围
			return math.MinInt32
		} else if sign*abs > math.MaxInt32 {
			return math.MaxInt32
		}
	}

	return abs * sign
}
