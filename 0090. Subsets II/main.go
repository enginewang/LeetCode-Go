package main

import (
	"fmt"
	"sort"
)

// 与78的区别在于，这里给出的nums可能有重复元素，而最后需要的结果不能包含重复
// 简单的想法是先全组合，然后去重，但是明显不是最佳方法
// 这里用了一个技巧，先排序，如果当前元素和上一个一样，就加一个判断，具体看代码
//func subsetsWithDup(nums []int) [][]int {
//	sort.Ints(nums)
//	result := make([][]int, 0, 1<<len(nums))
//	result = append(result, []int{})
//	var count []int
//	count = append(count, 1)
//	for i := 0; i < len(nums); i++ {
//		num := 0
//		if i > 0 && nums[i] == nums[i-1] {
//			for _, r := range result[len(result)-count[i]:] {
//				if len(r) > 0 && r[len(r)-1] == nums[i] {
//					newList := make([]int, len(r)+1)
//					copy(newList, r)
//					newList[len(newList)-1] = nums[i]
//					result = append(result, newList)
//					num++
//				}
//			}
//		} else {
//			for _, r := range result {
//				newList := make([]int, len(r)+1)
//				copy(newList, r)
//				newList[len(newList)-1] = nums[i]
//				result = append(result, newList)
//				num++
//			}
//		}
//		count = append(count, num)
//	}
//	return result
//}



//// 另一个方式是回溯法
//func subsetsWithDup(nums []int) [][]int {
//	sort.Ints(nums)
//	var ret [][]int
//	helper(nums, 0, []int{}, &ret)
//	return ret
//}
//
//func helper(nums []int, start int, temp []int, ret *[][]int) {
//	*ret = append(*ret, append([]int{}, temp...))
//	for i := start; i < len(nums); i++ {
//		if i > start && nums[i-1] == nums[i] {
//			continue
//		}
//		temp = append(temp, nums[i])
//		helper(nums, i+1, temp, ret)
//		temp = temp[:len(temp)-1]
//	}
//}



var result [][]int
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result = make([][]int, 0)
	dfs([]int{}, nums, 0)
	return result
}

func dfs(temp []int, nums []int, start int)  {
	tmp := make([]int, len(temp))
	copy(tmp, temp)
	result = append(result, tmp)
	for i := start; i < len(nums); i++{
		// 举例，假如是[1,1,2,2]，第一个1开始，后面所有的都可以，但是第二个1，直接pass，因为它
		// 开始temp的前面一个都做过一次了
		if i > start && nums[i] == nums[i-1]{
			continue
		}
		temp = append(temp, nums[i])
		dfs(temp, nums, i+1)
		temp = temp[:len(temp)-1]
	}
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 1, 2, 2}))
}