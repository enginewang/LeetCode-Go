package main

import "fmt"

var result [][]int
// 递归
//func permute(nums []int) [][]int {
//	result = make([][]int, 0)
//	visited := make([]int, len(nums))
//	dfs([]int{}, nums, visited)
//	return result
//}
//
//func dfs(temp []int, nums []int, visited []int)  {
//	if len(temp) == len(nums){
//		tmp := make([]int, len(temp))
//		copy(tmp, temp)
//		result = append(result, tmp)
//	}
//	for i := 0; i < len(nums); i++ {
//		if visited[i] == 1{
//			continue
//		}
//		temp = append(temp, nums[i])
//		visited[i] = 1
//		dfs(temp, nums, visited)
//		visited[i] = 0
//		temp = temp[:len(temp)-1]
//	}
//}

func permute(nums []int) [][]int {
	var path []int
	var result [][]int
	choiceDict := make(map[int] bool)
	for _, r := range nums {
		choiceDict[r] = true
	}
	backtrack(path, nums, choiceDict, &result)
	return result
}

func backtrack(path []int, nums []int, choiceDict map[int] bool, result *[][]int)  {
	if len(path) == len(choiceDict){
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
	}
	for _, choice := range nums{
		if choiceDict[choice] {
			path = append(path, choice)
			choiceDict[choice] = false
			backtrack(path, nums, choiceDict, result)
			choiceDict[choice] = true
			path = path[:len(path)-1]
		}
	}
}

// 12
// 12 21
// 123
// 312 132 123 321 231 213
func main() {
	fmt.Println(permute([]int{5,4,6,2}))
}
