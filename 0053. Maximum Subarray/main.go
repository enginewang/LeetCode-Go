package main

import "fmt"

// 动态规划
// 首先找关系
// dp(i)表示以第i个元素结尾的最大的子序和
// 如果dp(i-1)<0，那么dp(i)就是nums[i]，否则就是dp(i-1)+nums[i]
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	max := dp[0]
	for i := 1; i < len(nums); i++{
		if dp[i-1] < 0{
			dp[i] = nums[i]
		} else{
			dp[i] = dp[i-1] + nums[i]
		}
		if dp[i] > max{
			max = dp[i]
		}
	}
	return max
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
