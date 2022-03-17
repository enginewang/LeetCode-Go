package main

import (
	"fmt"
)

// 子数组之和的变形

func numberOfSubarrays(nums []int, k int) int {
	// 先构造S
	var S []int
	S = append(S, 0)
	for index := 1; index <= len(nums); index++ {
		i := nums[index-1] % 2
		S = append(S, S[index-1]+i)
	}
	sum := 0
	count := make([]int, len(S))
	for index := range S {
		count[S[index]] += 1
	}
	fmt.Println(count)
	fmt.Println(S)
	for j := range S {
		if S[j] >= k {
			sum += count[S[j]-k]
		}
	}
	return sum
}

func main() {
	fmt.Println(numberOfSubarrays([]int{1, 1, 2, 1, 1}, 3))
}
