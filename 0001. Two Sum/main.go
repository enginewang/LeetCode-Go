package main

/*
建立一个hashMap，遍历nums，假设nums[i]=num，存下map[target-num]=i，
并且如果有num的值存在于map中，可以取到之前的下标，就可以得到。
*/
import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, elem := range nums {
		_, ok := m[elem]
		if ok {
			return []int{m[elem], i}
		} else {
			m[target-elem] = i
		}
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
