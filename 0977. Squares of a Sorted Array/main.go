package main

import "fmt"

/*
给一个按非递减顺序排序的整数数组 `nums`，返回每个数字的平方组成的新数组，要求也按非递减顺序排序。

双指针

原数组是非递减的，所以先找到最小值所在的位置，然后两个指针往左右走。每次append小的数的平方到最终结果
*/

func sortedSquares(nums []int) []int {
	var result []int
	i, j := 0, 0
	for _, num := range nums {
		if num < 0 {
			i += 1
		} else {
			break
		}
	}
	j = i - 1
	for i < len(nums) || j >= 0 {
		if i >= len(nums) {
			result = append(result, nums[j]*nums[j])
			j--
		} else if j < 0 {
			result = append(result, nums[i]*nums[i])
			i++
		} else {
			if -nums[j] > nums[i] {
				result = append(result, nums[i]*nums[i])
				i++
			} else {
				result = append(result, nums[j]*nums[j])
				j--
			}
		}
	}
	return result
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}
