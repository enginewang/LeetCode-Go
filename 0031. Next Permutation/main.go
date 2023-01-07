package main

import (
	"fmt"
	"math"
	"sort"
)

func nextPermutation(nums []int) {
	find := false
	leftPos := 0
	// 首先从后往前找第一个变小的数
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			find = true
			leftPos = i - 1
			break
		}
	}
	if !find {
		for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	} else {
		fmt.Println(nums)
		// 找出右边最小的元素
		remainMin := math.MaxInt
		minPos := -1
		for j := leftPos + 1; j < len(nums); j++ {
			if nums[j] < remainMin && nums[j] > nums[leftPos] {
				remainMin = nums[j]
				minPos = j
			}
		}
		nums[leftPos], nums[minPos] = nums[minPos], nums[leftPos]
		fmt.Println(nums)
		sort.Ints(nums[leftPos+1:])
		fmt.Println(leftPos)
		fmt.Println(nums)
	}
}

func main() {
	nextPermutation([]int{2, 3, 1})
}

/*

1234
1243
1324
1342
1423
1432
2134
2143

*/
