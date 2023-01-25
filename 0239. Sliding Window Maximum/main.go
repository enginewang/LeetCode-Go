package main

import "fmt"

// hard
// window存当前最大的位置，第一个元素是最大的
func maxSlidingWindow(nums []int, k int) []int {
	if k == 0 || k > len(nums) {
		return make([]int, 0)
	}
	var result, windows []int
	for i := 0; i < len(nums); i++ {
		// 第一种情况，第一个元素移开了，需要退掉
		if i >= k && windows[0] <= i-k {
			windows = windows[1:len(windows)]
		}
		// 新来的元素比前面的大，从后往前一直退掉，如果比所有的都大，就只剩一个
		for len(windows) > 0 && nums[i] > nums[windows[len(windows)-1]] {
			windows = windows[0 : len(windows)-1]
		}
		// 移入该元素的位置
		windows = append(windows, i)
		if i >= k-1 {
			result = append(result, nums[windows[0]])
		}
	}
	return result
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, -1}, 1))
}
