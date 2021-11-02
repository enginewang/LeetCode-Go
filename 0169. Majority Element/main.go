package main

import "fmt"

func majorityElement(nums []int) int {
	cnt := 0
	result := 0
	for i := 0; i < len(nums); i++ {
		if cnt == 0 {
			result = nums[i]
			cnt++
		} else if nums[i] == result {
			cnt++
		} else {
			cnt--
		}
	}
	return result
}

func main() {
	fmt.Println(majorityElement([]int{2,3,3,3,2,2,2,2}))
}
