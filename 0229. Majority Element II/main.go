package main

import "fmt"

// 这里需要两个计数，因为最多也就两个
func majorityElement(nums []int) []int {
	var result []int
	if len(nums) == 1{
		return []int{nums[0]}
	}
	cand1 := nums[0]
	count1 := 0
	cand2 := nums[0]
	count2 := 0
	for i := 0; i < len(nums); i++{
		if nums[i] == cand1{
			count1++
			continue
		}
		if nums[i] == cand2{
			count2++
			continue
		}
		if count1 == 0{
			cand1 = nums[i]
			count1++
			continue
		}
		if count2 == 0{
			cand2 = nums[i]
			count2++
			continue
		}
		count1--
		count2--
	}
	count1, count2 = 0, 0
	// 只能说cand1和cand2可能超过1/3，还需要再计数一次
	for _, num := range nums{
		if num == cand1{
			count1++
		} else if num == cand2{
			count2++
		}
	}
	if count1 > len(nums)/3{
		result = append(result, cand1)
	}
	if count2 > len(nums)/3{
		result = append(result, cand2)
	}
	return result
}

func main() {
	fmt.Println(majorityElement([]int{4, 2, 1, 1}))
}
