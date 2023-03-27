package main

func maxProfit(prices []int, fee int) int {
	dp := make([][2]int, len(prices))
	// 第二位，0表示买入状态，1表示卖出状态
	dp[0][0], dp[0][1] = -prices[0], 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	}
	return dp[len(prices)-1][1]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
