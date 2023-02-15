package main

import "fmt"

var result [][]string
var dict map[string]bool

func partition(s string) [][]string {
	if len(s) == 1 {
		return [][]string{[]string{s}}
	}
	dict = make(map[string]bool)
	result = make([][]string, 0)
	backtrack([]int{0}, 0, s)
	return result
}

func backtrack(temp []int, start int, s string) {
	if start == len(s) {
		var r []string
		for k := 1; k < len(temp); k++ {
			r = append(r, s[temp[k-1]:temp[k]])
		}
		result = append(result, r)
	}
	for i := start + 1; i <= len(s); i++ {
		// 如果是回文子串，就进入递归
		if isReverseSame(s[start:i]) {
			temp = append(temp, i)
			backtrack(temp, i, s)
			temp = temp[:len(temp)-1]
		}
	}
}

func isReverseSame(s string) bool {
	if dict[s] {
		return dict[s]
	}
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			dict[s] = false
			return false
		}
	}
	dict[s] = true
	return true
}

func main() {
	fmt.Println(partition("aa"))
}
