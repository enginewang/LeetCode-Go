package main

// 移除所有数值等于 `val` 的元素，并返回移除后数组的新长度。
// 很简单，建立一个left指针，指向左边，遍历nums，如果读到了val，left就不动。否则就跟遍历一起往右移动。

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
