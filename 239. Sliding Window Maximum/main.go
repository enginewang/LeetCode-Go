package main

import "fmt"

// hard
// window存当前最大的位置，第一个元素是最大的

func maxSlidingWindow(nums []int, k int) []int {
	if k == 0 || k > len(nums) {
		return make([]int, 0)
	}
	result := make([]int, 0, len(nums)-k+1)
	windows := make([]int, 0, k)
	for i := 0; i < len(nums); i++ {
		// 第一种情况，第一个元素移开了，需要退掉
		if i >= k && windows[0] <= i-k{
			windows = windows[1:len(windows)]
		}

		for len(windows) > 0 && nums[i] > nums[windows[len(windows)-1]]{
			windows = windows[0:len(windows)-1]
		}
		// 移入该元素的位置
		windows = append(windows, i)
		if i >= k-1{
			if i == 3{
				fmt.Println("debug")
			}
			result = append(result, nums[windows[0]])
		}
	}
	return result
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
