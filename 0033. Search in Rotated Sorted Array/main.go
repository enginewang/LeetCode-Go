package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start := 0
	end := len(nums) - 1
	if nums[start] == target{
		return start
	}
	if nums[end] == target{
		return end
	}
	// 现在就是有序的二分搜索
	for start < end-1 {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		// 有序部分在前面
		if nums[mid] > nums[end]{
			if nums[mid] > target && nums[start] < target{
				end = mid
			} else {
				start = mid
			}
		} else {
			if nums[mid] < target && nums[end] > target{
				start = mid
			} else {
				end = mid
			}
		}
	}
	return -1
}

func main() {
	nums1 := []int{1, 3, 5}
	fmt.Println(search(nums1, 5))
	//nums2 := []int{4, 5, 6, 7, 0, 1, 2}
	//fmt.Println(search(nums2, 1))
}
