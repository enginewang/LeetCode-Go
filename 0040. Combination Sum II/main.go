package main

import (
	"fmt"
	"sort"
)

var result [][]int

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result = make([][]int, 0)
	backtrack([]int{}, 0, candidates, 0, target)
	return result
}

func backtrack(temp []int, sum int, nums []int, start int, target int) {
	if sum == target {
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		result = append(result, tmp)
		return
	}
	if sum > target {
		return
	}
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		temp = append(temp, nums[i])
		sum += nums[i]
		backtrack(temp, sum, nums, i+1, target)
		sum -= nums[i]
		temp = temp[:len(temp)-1]
	}
}

func main() {
	fmt.Println(combinationSum2([]int{1, 1, 2, 5, 6, 7, 10}, 8))
}
