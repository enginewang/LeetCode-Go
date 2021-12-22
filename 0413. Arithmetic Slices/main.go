package main

import "fmt"

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	cumulate := 0
	conti := 1
	num := 0
	for i := 2; i < len(nums); i++ {
		// 这个和之前是等差
		if nums[i] + nums[i-2] == 2 * nums[i-1] {
			cumulate += conti
			conti++
		} else {
			num += cumulate
			conti = 1
			cumulate = 0
		}
	}
	num += cumulate
	return num
}

func main() {
	fmt.Println(numberOfArithmeticSlices([]int{1,2,3,4}))
}
