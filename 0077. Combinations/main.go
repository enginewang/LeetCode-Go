package main

import "fmt"

func combine(n int, k int) [][]int {
	var result [][]int
	//cur := make([]int, 0, 5)
	backtracking(1, []int{}, &result, n, k)
	return result
}

func backtracking(start int, cur []int, result *[][]int, n int, k int) {
	if len(cur) == k {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*result = append(*result, tmp)
	}
	for i := start; i <= n; i++ {
		cur = append(cur, i)
		backtracking(i+1, cur, result, n, k)
		cur = cur[:len(cur)-1]
	}
}


func main() {
	fmt.Println(combine(10, 5))
}
