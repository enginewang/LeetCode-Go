package main

import "fmt"

func subarraySum(nums []int, k int) int {
	var preSumArray []int
	preSum := 0
	resultSum := 0
	for _, i := range nums {
		preSum += i
		preSumArray = append(preSumArray, preSum)
	}
	preSumArray = append([]int{0}, preSumArray...)
	auxMap := make(map[int] int)
	for i := 0; i < len(preSumArray); i++ {
		if _, ok := auxMap[preSumArray[i]]; ok{
			resultSum += auxMap[preSumArray[i]]
		}
		auxMap[preSumArray[i]+k] += 1
	}
	return resultSum
}

func main() {
	fmt.Println(subarraySum([]int{1,1,1,2,2,3,4,4,5,3,0,3,1,2}, 3))
}
