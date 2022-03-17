package main

import "fmt"


func maxArea(height []int) int {
	left, right, maxArea := 0, len(height)-1, 0
	for left < right{
		if height[left] < height[right] {
			area := height[left]*(right-left)
			if area > maxArea{
				maxArea = area
			}
			left++
		} else {
			area := height[right]*(right-left)
			if area > maxArea{
				maxArea = area
			}
			right--
		}
	}
	return maxArea
}


func main() {
	fmt.Println(maxArea([] int{1,8,6,2,5,4,8,3,7}))
}
