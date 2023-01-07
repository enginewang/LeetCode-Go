package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, targetet int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}
	sort.Ints(nums)
	var result [][]int
	// 跟三数之和一样，这里也是循环，找到i到最后的四个数字，且包括i
	for i := 0; i < len(nums)-3; i++ {
		// 去掉重复的，因为如果后面跟前面一样，那么前面肯定已经都找完了，再找都是重复的
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 这里因为是四个数字，需要再加一层跟上面类似的循环
		// 只有两个数字的和才能双指针
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right, target := j+1, len(nums)-1, targetet-nums[i]-nums[j]
			for left < right {
				sum := nums[left] + nums[right]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					right--
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					for left < right && nums[right] == nums[right+1] {
						right--
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Println(fourSum([]int{1, -2, -5, -4, -3, 3, 3, 5}, -11))
}
