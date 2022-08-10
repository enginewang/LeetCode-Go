package main

func minCostClimbingStairs(cost []int) int {
	// dp存的是爬到这层楼的最低花费
	dp := make([]int, len(cost)+1)
	dp[0], dp[1] = 0, 0
	for i := 2; i < len(cost)+1; i++ {
		dp[i] = min(dp[i-2] + cost[i-2], dp[i-1] + cost[i-1])
	}
	return dp[len(cost)]
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func main() {

}
