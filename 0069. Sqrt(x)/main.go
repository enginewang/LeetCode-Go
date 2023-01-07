package main

import "fmt"

func mySqrt(x int) int {
	left, right := 0, x
	// 在[left, right]中搜索平方根
	for left <= right {
		mid := (right - left) >> 1 + left
		mid2 := mid*mid
		if mid2 > x {
			right = mid - 1
		} else if mid2 < x {
			left = mid + 1
		} else {
			return mid
		}
	}
	if left < right {
		return left
	} else {
		return right
	}
}

func main() {
	fmt.Println(mySqrt(2))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(10))
	fmt.Println(mySqrt(100))
}
