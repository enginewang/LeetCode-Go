package main

import "fmt"

// 动态规划
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	fmt.Println(dp)
	return dp[n]
}

func main() {
	fmt.Println(numTrees(18))
}
