package main

func removeElement(nums []int, val int) int {
	left := 0
	for _, num := range nums {
		if num != val {
			nums[left] = num
			left += 1
		}
	}
	return left
}

func main() {
}
