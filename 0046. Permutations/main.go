package main

import "fmt"

var result [][]int
// 递归
func permute(nums []int) [][]int {
	result = make([][]int, 0)
	visited := make([]int, len(nums))
	dfs([]int{}, nums, visited)
	return result
}

func dfs(temp []int, nums []int, visited []int)  {
	if len(temp) == len(nums){
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		result = append(result, tmp)
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] == 1{
			continue
		}
		temp = append(temp, nums[i])
		visited[i] = 1
		dfs(temp, nums, visited)
		visited[i] = 0
		temp = temp[:len(temp)-1]
	}
}

// 12
// 12 21
// 123
// 312 132 123 321 231 213
func main() {
	fmt.Println(permute([]int{1,2,3}))
}
