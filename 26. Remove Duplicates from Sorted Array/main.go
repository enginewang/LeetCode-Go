package main

import "fmt"

// easy，没什么好说的

func removeDuplicates(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	p, q := 0, 1
	for ;q < len(nums); q++{
		if nums[q] > nums[q-1]{
			p++
			nums[p] = nums[q]
		}
	}
	fmt.Println(nums)
	return p+1
}

func main() {
	nums := []int {0,0,1,1,1,2,2,3,3,4}
	k := removeDuplicates(nums)
	fmt.Println(k)
}
