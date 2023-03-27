package main

import "fmt"

func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 4)
	}
	// 0 第一次持有该日股票
	// 1 第一次未持有该日股票
	// 2 第二次持有该日股票
	// 3 第二次未持有该日股票
	dp[0][0], dp[0][1], dp[0][2], dp[0][3] = -prices[0], 0, -prices[0], 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]-prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]+prices[i])
	}
	fmt.Println(dp)
	return dp[len(prices)-1][3]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
