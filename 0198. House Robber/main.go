package main

import "fmt"

// 动态规划
// 直接循环
// 对于第i项，考虑nums[i:]的能抢到的最大值
//func rob(nums []int) int {
//	memo := make([]int, len(nums)+2)
//	for i := 0; i < len(nums); i++ {
//		memo[i] = -1
//	}
//	return dp(0, nums, memo)
//}
//
//func dp(i int, nums []int, memo []int) int {
//	if i > len(nums) {
//		return 0
//	}
//	if memo[i] != -1 {
//		return memo[i]
//	} else {
//		res := max(nums[i]+dp(i+2, nums, memo), dp(i+1, nums, memo))
//		memo[i] = res
//		return res
//	}
//}
//
//func max(i int, j int) int {
//	if i > j {
//		return i
//	} else {
//		return j
//	}
//}

// 自底向上
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left, right := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		m := max(left + nums[i], right)
		left = right
		right = m
	}
	return right
}

func max(i int, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func main() {
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}
