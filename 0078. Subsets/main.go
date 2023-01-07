package main

import (
	"fmt"
	"sort"
)

//var res [][]int
//func subsets(nums []int) [][]int {
//	res = make([][]int, 0)
//	Dfs([]int{}, nums, 0)
//	return res
//}
//
//// temp表示当前的temp数组
//// start表示从第几个开始，前面的都不管了
//func Dfs(temp, nums []int, start int) {
//	tmp := make([]int, len(temp))
//	copy(tmp, temp)
//	res = append(res, tmp)
//	for i := start; i < len(nums); i++ {
//		temp = append(temp, nums[i])
//		Dfs(temp, nums, i+1)
//		temp = temp[:len(temp)-1]
//	}
//}

//func subsets(nums []int) [][]int {
//	result := make([][]int, 0, 1<<len(nums))
//	result = append(result, []int{})
//	for _, i := range nums{
//		for _, oldList := range result{
//			newList := make([]int, len(oldList)+1)
//			copy(newList, oldList)
//			newList[len(oldList)] = i
//			result = append(result, newList)
//		}
//	}
//	return result
//}

var result [][]int

func subsets(nums []int) [][]int {
	result = make([][]int, 0)
	sort.Ints(nums)
	backtrack([]int{}, nums, 0)
	return result
}

func backtrack(temp []int, nums []int, start int) {
	tmp := make([]int, len(temp))
	copy(tmp, temp)
	result = append(result, tmp)
	for i := start; i < len(nums); i++ {
		temp = append(temp, nums[i])
		backtrack(temp, nums, i+1)
		temp = temp[:len(temp)-1]
	}
}

func main() {
	fmt.Println(subsets([]int{1,2,3,4,5}))
}
