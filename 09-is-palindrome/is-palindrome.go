package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(10))
}

func isPalindrome(x int) bool {
	// 负数不可能为回文数
	// 数字为0结尾且不为0的数不可能未回文数
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	// 回文数只需要反转一半就知道结果
	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// 当偶数情况反转数等于x剩下的一般
	// 当奇数情况减少反转数末尾
	return x == revertedNumber || x == revertedNumber/10
}
