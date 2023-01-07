package main

import "fmt"

// 递归法
func numTrees(n int) int {
	memo := make(map[int]int)
	return helper(n, &memo)
}

func helper(n int, memo *map[int] int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if (*memo)[n] > 0 {
		return (*memo)[n]
	} else {
		result := 0
		for i := 0; i < n; i++ {
			result += helper(i, memo)*helper(n-1-i, memo)
		}
		(*memo)[n] = result
		return result
	}
}

func main() {
	fmt.Println(numTrees(18))
}
