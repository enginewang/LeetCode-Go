package main

import "fmt"

// 可以直接判断 str+str 中去除首尾元素之后，是否包含自身元素。如果包含。则表明存在重复子串。
// 采用kmp
func kmp(str string, pattern string) int {
	next := make([]int, len(pattern))
	next[0] = 0
	for i, j := 1, 0; i < len(pattern); i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = next[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		next[i] = j
	}
	//fmt.Println(next)
	for i, j := 0, 0; i < len(str); i++ {
		for j > 0 && str[i] != pattern[j] {
			j = next[j-1]
		}
		if str[i] == pattern[j] {
			j++
		}
		if j == len(pattern) {
			return i - len(pattern) + 1
		}
	}
	return -1
}

func repeatedSubstringPattern(s string) bool {
	newStr := s + s
	if kmp(newStr[1:len(newStr)-1], s) == -1 {
		return false
	}
	return true
}

func main() {
	fmt.Println(repeatedSubstringPattern("abcabcd"))
}
