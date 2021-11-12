package main

import "fmt"

// 动态规划
//func uniquePaths(m int, n int) int {
//	memo := make([][]int, m)
//	for i := 0; i < m; i++ {
//		memo[i] = make([]int, n)
//	}
//	return recursion(0, 0, m, n, memo)
//}
//
//func recursion(x int, y int, m int, n int, memo [][]int) int {
//	if x == m-1 || y == n-1 {
//		return 1
//	}
//	if memo[y][x] != 0{
//		return memo[y][x]
//	}
//	result := recursion(x+1, y, m, n, memo) + recursion(x, y+1, m, n, memo)
//	memo[y][x] = result
//	return result
//}

// 自底循环
// 只存一维
func uniquePaths(m int, n int) int {
	memo := make([]int, n)
	for i := 0; i < n; i++ {
		memo[i] = 1
	}
	for y := m - 2; y >= 0; y-- {
		for x := n - 2; x >= 0; x-- {
			memo[x] = memo[x] + memo[x+1]
		}
	}
	fmt.Println(memo)
	return memo[0]
}

func main() {
	fmt.Println(uniquePaths(3, 7))
}
