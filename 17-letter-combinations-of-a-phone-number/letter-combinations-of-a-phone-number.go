package main

import "fmt"

func main() {
	//fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations1("23"))
}

var phoneMap = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

var combinations []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		letters := phoneMap[digits[index]]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}
}

func letterCombinations1(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	dfs(digits, 0, "")
	return combinations
}

func dfs(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		for _, ch := range phoneMap[digits[index]] {
			dfs(digits, index+1, combination+string(ch))
		}
	}
}
