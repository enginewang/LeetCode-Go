package main

import (
	"fmt"
	"math"
)

// 迭代
func coinChange(coins []int, amount int) int {
	// dp[amount]表示在amount金额下最少需要的硬币个数
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	} else {
		return dp[amount]
	}
}

func min(i int, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func main() {
	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
}
