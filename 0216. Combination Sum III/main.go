package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	var result [][]int
	backtrack([]int{}, 1, 0, k, n, &result)
	return result
}

func backtrack(curPath []int, start int, sum int, count int, target int, result *[][]int) {
	if len(curPath) == count {
		if sum == target {
			tmp := make([]int, len(curPath))
			copy(tmp, curPath)
			*result = append(*result, tmp)
			return
		} else {
			return
		}
	}
	if sum >= target {
		return
	}
	for i := start; i < 10; i++ {
		curPath = append(curPath, i)
		sum += i
		backtrack(curPath, i+1, sum, count, target, result)
		sum -= i
		curPath = curPath[:len(curPath)-1]
	}
}

func main() {
	fmt.Println(combinationSum3(3, 9))
}
