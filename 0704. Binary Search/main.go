package main

import "fmt"

/*
二分查找

二分查找本身原理比较简单，是一种分治的策略，但是实际写代码的时候，边界问题容易出错。

搜索有序数组中target的位置，没找到则返回-1

首先要明确的是搜索的边界范围，是`[left, right]`还是`[left, right)`，比如下面两种都是对的。

推荐使用mid := left + (right-left)/2，而不是直接(left+right)/2，后者可能越界
*/

// 左闭右闭 [left, right]
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// 左闭右开 [left, right)
//func search(nums []int, target int) int {
//	left, right := 0, len(nums)
//	// 相等无意义
//	for left < right {
//		mid := left + (right-left)/2
//		if nums[mid] < target {
//			left = mid + 1
//		} else if nums[mid] > target {
//			right = mid
//		} else {
//			return mid
//		}
//	}
//	return -1
//}

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
}
