package main

import (
	"fmt"
	"math"
)

func jump(nums []int) int {
	memo := make([]int, len(nums))
	if len(nums) == 1{
		return 0
	}
	return dp(0, nums, memo)
}

// 最少的步数
func dp(i int, nums []int, memo []int) int {
	if nums[i]+i >= len(nums)-1 {
		return 1
	}
	if memo[i] > 0 {
		return memo[i]
	} else {
		min := math.MaxInt32
		for j := i + 1; j < i+nums[i]+1; j++ {
			minj := dp(j, nums, memo)
			if minj < min{
				min = minj
			}
		}
		memo[i] = min + 1
		return min + 1
	}
}

func main() {
	fmt.Println(jump([]int{2,3,1,1,4}))
}
