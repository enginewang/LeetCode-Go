package main

import "fmt"

var result [][]int

func combinationSum3(k int, n int) [][]int {
	result = make([][]int, 0)
	backtrack([]int{}, 0, 1, n, k)
	return result
}

func backtrack(temp []int, sum int, start int, target int, count int) {
	if len(temp) == count {
		if sum == target {
			tmp := make([]int, len(temp))
			copy(tmp, temp)
			result = append(result, tmp)
			return
		} else {
			return
		}
	}
	if sum > target {
		return
	}
	// 因为可以被无限次选取，temp表示从start开始到最后的结果
	for i := start; i < 10; i++ {
		temp = append(temp, i)
		sum += i
		backtrack(temp, sum, i+1, target, count)
		sum -= i
		temp = temp[:len(temp)-1]
	}
}

func main() {
	fmt.Println(combinationSum3(3, 9))
}
