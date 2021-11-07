package main

import "fmt"

// dfs
func allPathsSourceTarget(graph [][]int) [][]int {
	//visited := make([]bool, len(graph))
	var dfs func(int, []int)
	var result [][]int
	dfs = func(i int, path []int) {
		path = append(path, i)
		if i == len(graph)-1 {
			// 务必要这样，直接path有问题
			ans := make([]int, len(path))
			copy(ans, path)
			result = append(result, ans)
			return
		}
		for _, id := range graph[i] {
			dfs(id, path)
		}
	}
	dfs(0, []int{})
	return result
}

func main() {
	isConnected := [][]int{
		{4, 3, 1},
		{3, 2, 4},
		{3},
		{4},
		{},
	}
	fmt.Println(allPathsSourceTarget(isConnected))
}
