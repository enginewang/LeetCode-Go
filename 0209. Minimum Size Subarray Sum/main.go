package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	minLength := len(nums) + 1
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		if sum < target {
			continue
		}
		for sum >= target && left <= right {
			if right-left < minLength {
				minLength = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}
	if minLength == len(nums) + 1{
		return 0
	}
	return minLength
}

func main() {
	fmt.Println(minSubArrayLen(11, []int{1,1,1,1,1,1,1}))
}
