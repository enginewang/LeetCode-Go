package main

import (
	"math"
)

func minFallingPathSum(matrix [][]int) int {
	result := math.MaxInt
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i == 0 {
				dp[0][j] = matrix[0][j]
			} else if j == 0 {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
			} else if j == len(matrix[0])-1 {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + matrix[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
			}
			if i == len(matrix)-1 && dp[i][j] < result {
				result = dp[i][j]
			}
		}
	}
	return result
}

func min(intList ...int) int {
	result := math.MaxInt
	for _, i := range intList {
		if i < result {
			result = i
		}
	}
	return result
}

func main() {

}
