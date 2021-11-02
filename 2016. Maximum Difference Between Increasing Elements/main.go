package main

import "fmt"

// 这个和121买股票完全一样

func maximumDifference(nums []int) int {
	prevMin := 0
	max := -1
	for i := 1; i < len(nums); i++{
		if nums[i] - nums[prevMin] >  max{
			max = nums[i] - nums[prevMin]
		}
		if nums[i] < nums[prevMin]{
			prevMin = i
		}
	}
	if max > 0{
		return max
	} else {
		return -1
	}
}

func main() {
	fmt.Println(maximumDifference([]int{1, 5, 2, 10}))
}

// [7, 9, 1, 2 ,3 , 100]
