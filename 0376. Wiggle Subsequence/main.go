package main

import "fmt"

func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	count := 1
	last := nums[0]
	lastIncr := 1
	for i := 1; i < len(nums); i++ {
		if (nums[i]-last)*lastIncr >= 0 {
			last = nums[i]
		} else {
			count++
			last = nums[i]
			lastIncr = -lastIncr
		}
		fmt.Println(last, lastIncr, count)
	}
	return count
}

func main() {
	wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8})
}
