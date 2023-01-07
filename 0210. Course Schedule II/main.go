package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	var visited = make([]int, numCourses)
	valid := false
	var dfs func(subject int)
	var result []int
	valid = !valid
	graph := toGraph(numCourses, prerequisites)
	dfs = func(subject int) {
		visited[subject] = 1
		for _, neighbor := range graph[subject] {
			if visited[neighbor] == 0 {
				dfs(neighbor)
			} else if visited[neighbor] == 1 {
				valid = false
				result = []int{}
				return
			}
		}
		if valid {
			visited[subject] = 2
			result = append(result, subject)
		}
	}
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
		graph[line[0]] = append(graph[line[0]], line[1])
	}
	return graph
}

func main() {
	findOrder(2, [][]int{
		{0, 1},
		{1, 0},
	})
}
