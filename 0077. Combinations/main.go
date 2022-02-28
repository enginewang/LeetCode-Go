package main

import "fmt"

var result [][]int

func combine(n int, k int) [][]int {
	backtracking(1, []int{}, n, k)
	return result
}

func backtracking(start int, cur []int, n int, k int) {
	if len(cur) == k {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		result = append(result, tmp)
	}
	for i := start; i <= n; i++ {
		cur = append(cur, i)
		backtracking(i+1, cur, n, k)
		cur = cur[:len(cur)-1]
	}
}


func main() {
	fmt.Println(combine(10, 5))
}
