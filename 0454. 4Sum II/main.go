package main

import (
	"fmt"
)

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// 构造一个map，k是前两个数组可能组成的一个和，v是次数
	m := make(map[int]int)
	result := 0
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			m[v1+v2] += 1
		}
	}
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			result += m[-v3-v4]
		}
	}
	return result
}

func main() {
	fmt.Println(fourSumCount([]int{-1, -1}, []int{-1, 1}, []int{-1, 1}, []int{1, -1}))
}
