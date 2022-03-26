package main

import (
	"fmt"
	"sort"
)


func threeSum(nums []int) [][]int {
	var results [][]int
	// 先对于所有的元素进行排序
	sort.Ints(nums)
	// 这里每一次for循环是从nums[i,length-1]中找到包含i的所有三元组
	// 1. 所以到了i-2即可，三元组不可能不包含前length-2个元素的任何一个
	// 2. 下面的防止repeat同理，假设前一个元素找出了包含它的所有三元组，后一个元素如果值一样，那后一个直接跳过
	// 3. 下面的left,right之从后面取也是这样原因，之前的通通不考虑了
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			return results
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue//To prevent the repeat
		}
		target, left, right := -nums[i], i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				results = append(results, []int{nums[i], nums[left], nums[right]})
				// 左右先互相靠拢，因为必然包含了i，如果还包含left对应的值，就必然重复
				left++
				right--
				// 这边也是一样，左右如果有相同的就跳，见上一条注释
				// 通过这些方式，生成的结果必然不包含重复的值，最后不需要set操作
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > target {
				right--
			} else if sum < target {
				left++
			}
		}
	}
	return results
}

//func threeSum(nums []int) [][]int {
//	var result [][]int
//	if len(nums) < 3{
//		return result
//	}
//	sort.Ints(nums)
//	for i := 0; i < len(nums)-2; i++ {
//		if i >= 1 && nums[i] == nums[i-1] {
//			continue
//		}
//		left, right, target := i+1, len(nums)-1, -nums[i]
//		for left < right {
//			if nums[left]+nums[right] == target {
//				result = append(result, []int{nums[i], nums[left], nums[right]})
//				for left < right && nums[left] == nums[left+1] && nums[right] == nums[right-1] {
//					left++
//					right--
//				}
//				left++
//				right--
//			} else if nums[left]+nums[right] < target {
//				left++
//			} else {
//				right--
//			}
//		}
//	}
//	return result
//}


// -4 -1 -1 0 1 2
func main() {
	fmt.Println(threeSum([]int{-2, 0, 0, 2, 2}))
}
