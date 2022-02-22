package main

import "fmt"

// dfs
//func allPathsSourceTarget(graph [][]int) [][]int {
//	var dfs func(int, []int)
//	var result [][]int
//	dfs = func(i int, path []int) {
//		path = append(path, i)
//		if i == len(graph)-1 {
//			// 务必要这样，直接path有问题
//			ans := make([]int, len(path))
//			copy(ans, path)
//			result = append(result, ans)
//			return
//		}
//		for _, id := range graph[i] {
//			dfs(id, path)
//		}
//	}
//	dfs(0, []int{})
//	return result
//}


// 图的遍历
func allPathsSourceTarget(graph [][]int) [][]int {
	var result [][]int
	var path []int
	traverse(graph, path, 0, &result)
	return result
}

func traverse(graph [][]int, path []int, c int, result *[][]int)  {
	path = append(path, c)
	// 当前结点是最后一个，满足条件
	if c == len(graph) - 1{
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
	}
	for _, child := range graph[c] {
		traverse(graph, path, child, result)
	}
	path = path[:len(path)-1]
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
