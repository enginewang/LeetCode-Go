package main

import "fmt"

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	left := 0
	right := len(nums) - 1
	for left < right-1 {
		mid := (left + right) / 2
		if judgeMax(nums, mid) {
			return mid
		}
		if nums[mid] < nums[right] {
			left = mid
		} else if nums[mid] < nums[left]{
			right = mid
		} else {
			left++
			right--
		}
	}
	if judgeMax(nums, left) {
		return left
	}
	if judgeMax(nums, right) {
		return right
	}

	return -1
}

func judgeMax(nums []int, i int) bool {
	if i == 0 {
		return nums[i] > nums[i+1]
	}
	if i == len(nums)-1 {
		return nums[i] > nums[i-1]
	}
	if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
		return true
	}
	return false
}

func main() {
	fmt.Println(findPeakElement([]int{1,6,5,4,3,2,1}))
}
