package main

import "fmt"

// 动态规划，memo[i][j]存的是s[i,j]是否是回文串
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	ans := string(s[0])
	for r := 0; r < len(s); r++ {
		dp[r][r] = true
		for l := 0; l < r; l++ {
			if s[l] == s[r] && (l+1 > r-1 || dp[l+1][r-1]) {
				dp[l][r] = true
				if r-l+1 > len(ans) {
					ans = s[l : r+1]
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(longestPalindrome("ac"))
}
