package main

import "fmt"

func searchInsert(nums []int, target int) int {
	// 在[left, right]里寻找
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right - left) >> 1 + left
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return left
}

func main() {
	fmt.Println(searchInsert([]int{1,3,5,6}, 5))
	fmt.Println(searchInsert([]int{1,3,5,6}, 2))
	fmt.Println(searchInsert([]int{1,3,5,6}, 7))
	fmt.Println(searchInsert([]int{1,3,5,6}, 0))
	fmt.Println(searchInsert([]int{1}, 0))
}
