package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	memo := make([][]int, len(obstacleGrid))
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, len(obstacleGrid[0]))
	}
	for row := 0; row < len(obstacleGrid); row++ {
		for col := 0; col < len(obstacleGrid[0]); col++ {
			if obstacleGrid[row][col] == 1 {
				memo[row][col] = 0
			} else if row == 0 && col == 0 {
				memo[row][col] = 1
			} else if row == 0 {
				memo[row][col] = memo[row][col-1]
			} else if col == 0 {
				memo[row][col] = memo[row-1][col]
			} else {
				memo[row][col] = memo[row-1][col] + memo[row][col-1]
			}
		}
	}
	//fmt.Println(memo)
	return memo[len(memo)-1][len(memo[0])-1]
}

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}))
}
