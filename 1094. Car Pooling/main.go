package main

import "fmt"

func carPooling(trips [][]int, capacity int) bool {
	maxLength := 0
	for _, trip := range trips{
		if trip[2] > maxLength{
			maxLength = trip[2]
		}
	}
	diffArray := make([]int, maxLength+1)
	for _, trip := range trips{
		numP := trip[0]
		from := trip[1]
		to := trip[2]
		diffArray[from] += numP
		diffArray[to] -= numP
	}
	cur := 0
	for _, i := range diffArray{
		cur += i
		if cur > capacity{
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(carPooling([][]int { {2,1,5}, {3,3,7}}, 5))
}
