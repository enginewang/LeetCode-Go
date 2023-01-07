package main

import (
	"fmt"
	"math"
)

// 动态规划，dp[i][j]表式骑士到从(i,j)到终点需要的最少的生命
func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}
	for i, _ := range dp {
		dp[i][n] = math.MaxInt
	}
	for j, _ := range dp[0] {
		dp[m][j] = math.MaxInt
	}
	if dungeon[m-1][n-1] < 0 {
		dp[m-1][n-1] = 1 - dungeon[m-1][n-1]
	} else {
		dp[m-1][n-1] = 1
	}
	//printMatrix(dp)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				continue
			}
			minLife := min(dp[i+1][j], dp[i][j+1])
			if minLife <= dungeon[i][j] {
				dp[i][j] = 1
			} else {
				dp[i][j] = minLife - dungeon[i][j]
			}
		}
	}
	printMatrix(dp)
	return dp[0][0]
}

func min(i int, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func printMatrix(m [][]int) {
	for _, row := range m {
		fmt.Println(row)
	}
	fmt.Println("---------")
}

func main() {
	m := [][]int{
		{0, 0, 0},
		{1, 1, -1},
	}
	printMatrix(m)
	calculateMinimumHP(m)
}
