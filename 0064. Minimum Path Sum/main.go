package main

import "math"

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid)+1)
	for i := 0; i <= len(grid); i++ {
		dp[i] = make([]int, len(grid[0])+1)
	}
	for i := 0; i <= len(grid); i++ {
		dp[i][0] = math.MaxInt32
	}
	for j := 0; j <= len(grid[0]); j++{
		dp[0][j] = math.MaxInt32
	}
	for i := 1; i <= len(grid); i++ {
		for j := 1; j <= len(grid[0]); j++{
			if i == 1 && j == 1{
				dp[i][j] = grid[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i-1][j-1]
			}
		}
	}
	return dp[len(grid)][len(grid[0])]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

}
