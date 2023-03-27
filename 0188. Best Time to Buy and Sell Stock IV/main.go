package main

func maxProfit(k int, prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2*k)
	}
	// 2*i 第i次持有股票
	// 2*i+1 第i次未持有股票
	for i := 0; i < k; i++ {
		dp[0][2*i] = -prices[0]
		dp[0][2*i+1] = 0
	}
	for i := 1; i < len(prices); i++ {
		for j := 0; j < 2*k; j++ {
			if j == 0 {
				dp[i][j] = max(dp[i-1][j], -prices[i])
			} else if j%2 == 0 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]-prices[i])
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]+prices[i])
			}
		}
	}
	return dp[len(prices)-1][2*k-1]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
