package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(robRange(0, len(nums)-2, nums), robRange(1, len(nums)-1, nums))
}

// 计算从start到end区间内的最优解
func robRange(start int, end int, nums []int) int {
	if end == start {
		return nums[start]
	}
	m := 0
	left, right := nums[start], max(nums[start], nums[start+1])
	for i := start + 2; i <= end; i++ {
		m = max(left+nums[i], right)
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
	fmt.Println(rob([]int{1, 2, 3, 1}))
}
