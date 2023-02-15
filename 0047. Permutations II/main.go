package main

import (
	"fmt"
	"sort"
)

var result [][]int

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	result = make([][]int, 0)
	visited := make([]int, len(nums))
	dfs([]int{}, nums, visited)
	return result
}

func dfs(temp []int, nums []int, visited []int) {
	if len(temp) == len(nums) {
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		result = append(result, tmp)
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] == 1 {
			continue
		}
		// 为了去掉可能重复的情况，如果之前访问过i-1，那么下一个i直接pass
		// 如果是说之前的没访问过，这里pass也正确，最后改成==0
		if i > 0 && nums[i] == nums[i-1] && visited[i-1] == 1 {
			continue
		}
		visited[i] = 1
		temp = append(temp, nums[i])
		dfs(temp, nums, visited)
		visited[i] = 0
		temp = temp[:len(temp)-1]
	}
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2, 3}))
}
