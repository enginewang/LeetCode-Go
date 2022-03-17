package main

import "fmt"

// 和26题类似，没什么好说的
func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	p, q := 0, 0
	for ; q < len(nums); q++ {
		if nums[q] != 0 {
			nums[p] = nums[q]
			p++
		}
	}
	for ; p < len(nums); p++{
		nums[p] = 0
	}
}

func main() {
	nums := []int{0,1,0,3,12}
	moveZeroes(nums)
	fmt.Println(nums)
}
