package main

import "fmt"

// 本质上就是找图有没有环
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 遍历一遍图，加一个visited，如果访问到了之前的就返回false
	visited := make([]int, numCourses)
	graph := toGraph(numCourses, prerequisites)
	var dfs func(u int)
	var result bool = true
	dfs = func(subject int) {
		visited[subject] = 1
		for _, neighbors := range graph[subject] {
			if visited[neighbors] == 1 {
				result = false
				return
			} else if visited[neighbors] == 0 {
				dfs(neighbors)
			}
		}
		visited[subject] = 2
	}
	//fmt.Println(graph)
	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return result
}

func toGraph(numCourses int, prerequisites [][]int) [][]int {
	graph := make([][]int, numCourses)
	for _, line := range prerequisites {
		graph[line[1]] = append(graph[line[1]], line[0])
	}
	return graph
}

func main() {
	fmt.Println(canFinish(5, [][]int{
		{1, 4},
		{2, 4},
		{3, 1},
		{3, 2},
	}))
}
