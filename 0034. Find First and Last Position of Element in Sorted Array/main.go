package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0{
		return []int{-1, -1}
	}
	// 二分查找
	start := 0
	end := len(nums)-1
	findOne := -1
	if nums[start] == target{
		findOne = start
	} else if nums[end] == target{
		findOne = end
	} else {
		for start < end-1 {
			mid := (start+end)/2
			if nums[mid] == target{
				findOne = mid
				break
			}
			if nums[mid] < target{
				start = mid
				continue
			} else {
				end = mid
				continue
			}
		}
	}
	if findOne == -1{
		return []int{-1, -1}
	} else {
		firstId, lastId := findOne, findOne
		for firstId >= 0 && nums[firstId] == target{
			firstId--
		}
		for lastId <= len(nums)-1 && nums[lastId] == target{
			lastId++
		}
		return []int{firstId+1, lastId-1}
	}
}

func main() {
	nums := []int{1}
	fmt.Println(searchRange(nums, 1))
}
