package main

import "fmt"

func totalFruit(fruits []int) int {
	left, right, maxAmount := 0, 0, 0
	fruitMap := make(map[int]int)
	for ; right < len(fruits); right++ {
		fruitMap[fruits[right]] += 1
		// 右移left
		for len(fruitMap) > 2 {
			fruitMap[fruits[left]] -= 1
			if fruitMap[fruits[left]] == 0{
				delete(fruitMap, fruits[left])
			}
			left += 1
		}
		if len(fruitMap) == 2 && maxAmount < right-left+1 {
			maxAmount = right - left + 1
		}
	}
	if right-left > maxAmount{
		maxAmount = right-left
	}
	return maxAmount
}

func main() {
	fruits := []int{3,3,3,1,2,1,1,2,3,3,4}
	fmt.Println(totalFruit(fruits))
}
