package main

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) int {
	left := 0
	curMul := 1
	result := 0
	for right := 0; right < len(nums); right++ {
		// 满足条件
		curMul *= nums[right]
		for curMul >= k && left <= right{
			curMul /= nums[left]
			left++
		}
		if left <= right {
			// 因为这里会新生成right - left + 1个以right结尾的新字符串
			result += right - left + 1
		}
	}
	return result
}

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10,5,2,6}, 100))
}
