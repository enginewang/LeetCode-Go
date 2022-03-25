package main

import "fmt"

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
	j = i-1
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
	fmt.Println(sortedSquares([]int{-4,-1,0,3,10}))
}
