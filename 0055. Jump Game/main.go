package main

import "fmt"

// 动态规划，会超时
//func canJump(nums []int) bool {
//	memo := make([]bool, len(nums))
//	memo[len(nums)-1] = true
//	result := dp(0, nums, memo)
//	fmt.Println(memo)
//	return result
//}
//
//func dp(i int, nums []int, memo []bool) bool {
//	if memo[i] == true {
//		return true
//	}
//	for j := i + 1; j < i+nums[i]+1; j++ {
//		if dp(j, nums, memo) == true{
//			memo[i] = true
//			return true
//		}
//	}
//	memo[i] = false
//	return false
//}

// 从后向前循环，不会超时
func canJump(nums []int) bool {
	firstTruePos := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < firstTruePos-i {
			continue
		} else {
			firstTruePos = i
		}
	}
	if firstTruePos != 0 {
		return false
	} else {
		return true
	}
}

// 贪心，每次维护最远的距离
//func canJump(nums []int) bool {
//	maxPos := 0
//	for i := 0; i < len(nums); i++ {
//		if i > maxPos {
//			return false
//		}
//		if i + nums[i] > maxPos{
//			maxPos = i + nums[i]
//		}
//	}
//	return true
//}

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
}
