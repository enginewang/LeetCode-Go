package main

import "fmt"

func isBipartite(graph [][]int) bool {
	visited := make([]bool, len(graph))
	color := make([]bool, len(graph))
	ok := true
	for i := 0; i < len(graph); i++ {
		if !visited[i]{
			traverse(graph, visited, color, i, &ok)
		}
	}
	return ok
}

func traverse(graph [][]int, visited []bool, color []bool, c int, ok *bool) {
	if !*ok {
		return
	}
	visited[c] = true
	for _, child := range graph[c]{
		if !visited[child] {
			color[child] = !color[c]
			traverse(graph, visited, color, child, ok)
		} else {
			if color[c] == color[child]{
				*ok = false
			}
		}
	}
}

func main() {
	fmt.Println(isBipartite([][]int{
		{1,2,3},
		{0,2},
		{0,1,3},
		{0,2},
	}))
}
