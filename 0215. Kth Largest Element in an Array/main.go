package main

// 方法一，采用快速排序的思想
func findKthLargest(nums []int, k int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		p := partition(nums, left, right)
		if p == len(nums)-k {
			return nums[p]
		} else if p < len(nums)-k {
			left = p + 1
		} else {
			right = p - 1
		}
	}
	return 0
}

func partition(nums []int, left int, right int) int {
	pivot := nums[right]
	i := left
	for j := left; j <= right; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

func main() {

}
