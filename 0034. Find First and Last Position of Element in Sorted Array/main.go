package main

import "fmt"

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)>>1 + left
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	findOne := binarySearch(nums, target)
	if findOne == -1 {
		return []int{-1, -1}
	} else {
		firstId, lastId := findOne, findOne
		for firstId >= 0 && nums[firstId] == target {
			firstId--
		}
		for lastId <= len(nums)-1 && nums[lastId] == target {
			lastId++
		}
		return []int{firstId + 1, lastId - 1}
	}
}

func main() {
	nums := []int{1}
	fmt.Println(searchRange(nums, 1))
}
