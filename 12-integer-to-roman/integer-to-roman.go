package main

import "fmt"

func main() {
	fmt.Println(romanToInt("I"))
}

func romanToInt(s string) int {
	var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	l := len(s)
	ans := 0
	for i := range s {
		// 如果当前符号所表示的大小小于后面符号时，则为倒挂情况
		if i < l-1 && symbolValues[s[i]] < symbolValues[s[i+1]] {
			ans -= symbolValues[s[i]]
		} else {
			ans += symbolValues[s[i]]
		}
	}
	return ans
}
