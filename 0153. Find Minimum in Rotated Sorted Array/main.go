package main

import "fmt"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right-1{
		mid := (left + right) / 2
		if nums[mid] < nums[left]{
			right = mid
		} else if nums[mid] > nums[right]{
			left = mid
		} else {
			return nums[left]
		}
	}
	if nums[left] > nums[right]{
		return nums[right]
	} else {
		return nums[left]
	}
}

func main() {
	fmt.Println(findMin([]int{1}))
}
