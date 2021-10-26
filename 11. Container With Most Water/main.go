package main

import "fmt"

func maxArea(height []int) int {
	max := 0
	left := 0
	right := len(height)-1
	for left < right{
		if height[left] < height[right] {
			curArea := height[left] * (right - left)
			if curArea > max {
				max = curArea
			}
			left++
		} else {
			curArea := height[right] * (right - left)
			if curArea > max {
				max = curArea
			}
			right--
		}
	}
	return max
}

func main() {
	fmt.Println(maxArea([] int{1,2,1}))
}
