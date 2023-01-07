package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	// dp[i][j]的含义是word1[:i]和word2[:j]的最小编辑距离
	dp := make([][]int, len(word1)+1)
	for row, _ := range dp {
		dp[row] = make([]int, len(word2)+1)
	}
	for i := 0; i <= len(word1); i++ {
		dp[i][0] = i
	}
	for j := 0; j <= len(word2); j++ {
		dp[0][j] = j
	}
	for i := 0; i < len(word1); i++ {
		for j := 0; j < len(word2); j++ {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min(dp[i+1][j], dp[i][j+1], dp[i][j]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func min(vals ...int) int {
	ans := vals[0]
	for _, v := range vals {
		if ans > v {
			ans = v
		}
	}
	return ans
}

func main() {
	fmt.Println(minDistance("horssdfasdfasdfwefewfe", "rosdfgdhykjyiukgjutn"))
}
