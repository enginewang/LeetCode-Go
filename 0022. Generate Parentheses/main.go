package main

import "fmt"


// 左括号的个数，右括号的个数，当前的string
func recursion(left int, right int, n int, s string, result *[]string) {
	if len(s) == n {
		*result = append(*result, s)
	}
	// 只要左括号少于一半，就可以加
	if left < n/2 {
		recursion(left+1, right, n, s+"(", result)
	}
	// 如果左比右多，就可以加右括号
	if left > right {
		recursion(left, right+1, n, s+")", result)
	}
}

func generateParenthesis(n int) []string {
	var result []string
	recursion(0, 0, 2*n, "", &result)
	return result
}

func main() {
	fmt.Println(generateParenthesis(3))
}
