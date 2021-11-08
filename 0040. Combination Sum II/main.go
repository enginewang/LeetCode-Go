package main

import (
	"fmt"
	"sort"
)

var result [][]int
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result = make([][]int, 0)
	dfs([]int{}, 0, candidates, 0, target)
	return result
}

func dfs(temp []int, sum int, nums []int, start int, target int)  {
	if sum == target {
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		result = append(result, tmp)
	}
	if sum > target{
		return
	}
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1]{
			continue
		}
		sum += nums[i]
		temp = append(temp, nums[i])
		dfs(temp, sum, nums, i+1, target)
		sum -= nums[i]
		temp = temp[:len(temp)-1]
	}
}

func main() {
	fmt.Println(combinationSum2([]int{1,1,2,5,6,7,10}, 8))
}
