package main

import "fmt"

func quickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}
	q := partition(nums, left, right)
	quickSort(nums, left, q-1)
	quickSort(nums, q+1, right)
}

func partition(nums []int, left int, right int) int {
	pivot := nums[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] < pivot {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}
	nums[right], nums[i] = nums[i], nums[right]
	return i
}

func main() {
	nums := []int{1, 5, 3, 2, 6, 4, 9, 0}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
