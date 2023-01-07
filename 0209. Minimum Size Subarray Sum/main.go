package main

import (
	"fmt"
	"math"
)

/*
给定一个含有n个正整数的数组和一个正整数target，找出该数组中满足其和≥target的长度最小的 连续子数组并返回其长度

```
输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。

还是双指针，
*/

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	minLength := math.MaxInt32
	// 大循环
	for start, end := 0, 0; end < len(nums); end++ {
		sum += nums[end]
		// 如果sum >= target，更新答案，并让start往右走
		for sum >= target {
			if end-start+1 < minLength {
				minLength = end - start + 1
			}
			sum -= nums[start]
			start += 1
		}
	}
	if minLength == math.MaxInt32 {
		return 0
	}
	return minLength
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}
