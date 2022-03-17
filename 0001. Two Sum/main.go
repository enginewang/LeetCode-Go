package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int] int)
	for i, elem := range nums{
		_, ok := m[elem]
		if ok {
			return []int{m[elem], i}
		} else {
			m[target - elem] = i
		}
	}
	return nil
}


func main() {
	fmt.Println(twoSum([]int{3,2,4}, 6))
}
