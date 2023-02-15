package main

import (
	"fmt"
)

var result [][]int

func findSubsequences(nums []int) [][]int {
	result = make([][]int, 0)
	backtrack([]int{}, 0, nums)
	return result
}

func backtrack(temp []int, start int, nums []int) {
	if len(temp) > 1 {
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		result = append(result, tmp)
	}
	// 建立一个used字典，对同层元素进行去重
	used := make(map[int]bool, len(nums))
	for i := start; i < len(nums); i++ {
		// 每次同一层的时候才去重，比如[1 2 1 1]，对于一个元素开始，遍历后面的元素
		// 如果发现之前某个值有遍历过，那么后面的值直接跳过，用这个used来去重
		// 每一层重新建立一个used
		if used[nums[i]] {
			continue
		}
		if len(temp) == 0 || nums[i] >= temp[len(temp)-1] {
			temp = append(temp, nums[i])
			used[nums[i]] = true
			backtrack(temp, i+1, nums)
			temp = temp[:len(temp)-1]
		}
	}
}

func main() {
	fmt.Println(findSubsequences([]int{1, 2, 1, 1, 1, 1, 1}))
}
