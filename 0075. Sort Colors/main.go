package main

import "fmt"

func sortColors(nums []int) {
	countRed, countBlue := 0, 0
	for _, num := range nums {
		if num == 0 {
			countRed += 1
		}
		if num == 2 {
			countBlue += 1
		}
	}
	for i := 0; i < len(nums); i++ {
		if i < countRed {
			nums[i] = 0
		} else if i > len(nums)-countBlue-1 {
			nums[i] = 2
		} else {
			nums[i] = 1
		}
	}
}

func main() {
	n := []int{2, 0, 2, 1, 1, 0}
	sortColors(n)
	fmt.Println(n)
}
