package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	if len(nums) == 1{
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxResult := dp[0]
	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		if dp[i-1]>0{
			cur += dp[i-1]
		}
		dp[i] = cur
		if cur > maxResult{
			maxResult = cur
		}
	}
	fmt.Println(dp)
	return maxResult
}

func main() {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}
