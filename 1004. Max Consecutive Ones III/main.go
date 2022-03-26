package main

import "fmt"

func longestOnes(nums []int, k int) int {
	left, right := 0, 0
	maxNum := 0
	for ; right < len(nums); right++ {
		if nums[right] == 1 || k > 0 {
			if nums[right] == 0 {
				k--
			}
		} else {
			if right-left > maxNum {
				maxNum = right - left
				fmt.Println(left, right)
			}
			k--
			for k < 0 {
				if nums[left] == 0 {
					k++
				}
				left++
			}
		}
	}
	if right - left > maxNum {
		maxNum = right - left
	}
	return maxNum
}

func main() {
	fmt.Println(longestOnes([]int{0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1}, 3))
}
