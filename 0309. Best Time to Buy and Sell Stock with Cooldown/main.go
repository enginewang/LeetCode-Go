package main

func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 4)
	}
	dp[0][0], dp[0][1], dp[0][2], dp[0][3] = -prices[0], 0, 0, 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], max(dp[i-1][2], dp[i-1][3])-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = dp[i-1][1]
		dp[i][3] = max(dp[i-1][3], dp[i-1][2])
	}
	return max(max(dp[len(prices)-1][1], dp[len(prices)-1][2]), dp[len(prices)-1][3])
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
