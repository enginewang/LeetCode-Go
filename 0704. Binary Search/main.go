package main

import "fmt"

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

//func search(nums []int, target int) int {
//	// 变为[left, right)
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
