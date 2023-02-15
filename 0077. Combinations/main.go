package main

import "fmt"

func combine(n int, k int) [][]int {
	var result [][]int
	backtracking(1, []int{}, n, k, &result)
	return result
}

func backtracking(start int, cur []int, n int, k int, result *[][]int) {
	if len(cur) == k {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*result = append(*result, tmp)
	}
	for i := start; i <= n; i++ {
		// 不可能满足要求，剪枝
		if n-i+1 < k-len(cur) {
			break
		}
		cur = append(cur, i)
		backtracking(i+1, cur, n, k, result)
		cur = cur[:len(cur)-1]
	}
}

func main() {
	fmt.Println(combine(10, 5))
}
