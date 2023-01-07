package main

import "fmt"

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
	return nums
}

// 在[left, right]之间排好序
func quickSort(nums []int, left int, right int) {
	if left > right {
		return
	}
	pivot, i, j := nums[left], left, right
	// 找到pivot在其中的正确位置
	for i < j {
		for i < j && nums[j] >= pivot {
			j -= 1
		}
		for i < j && nums[i] <= pivot {
			i += 1
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}

func main() {
	sortArray([]int{5, 1, 1, 2, 0, 0})
}
