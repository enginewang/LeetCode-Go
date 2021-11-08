package main

import "fmt"

var result [][]int

func combinationSum(candidates []int, target int) [][]int {
	result = make([][]int, 0)
	dfs([]int{}, 0, candidates, 0, target)
	return result
}

func dfs(temp []int, sum int, nums []int, start int, target int) {
	if sum == target {
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		result = append(result, tmp)
		return
	}
	if sum > target {
		return
	}
	// 因为可以被无限次选取，temp表示从start开始到最后的结果
	for i := start; i < len(nums); i++ {
		temp = append(temp, nums[i])
		sum += nums[i]
		dfs(temp, sum, nums, i, target)
		sum -= nums[i]
		temp = temp[:len(temp)-1]
	}
}

func main() {
	fmt.Println(combinationSum([]int{2}, 6))
}
