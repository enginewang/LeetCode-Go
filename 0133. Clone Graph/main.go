package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	return dfs(node, &visited)
}

func dfs(node *Node, visited *map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	if _, ok := (*visited)[node]; ok {
		return (*visited)[node]
	}
	clone := &Node{
		Val:       node.Val,
		Neighbors: []*Node{},
	}
	(*visited)[node] = clone
	for _, neighbor := range node.Neighbors {
		// 如果有邻居没访问就去访问
		(*visited)[node].Neighbors = append((*visited)[node].Neighbors, dfs(neighbor, visited))
	}
	return clone
}

func main() {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node1.Neighbors = append(node1.Neighbors, node2)
	node1.Neighbors = append(node1.Neighbors, node4)
	node2.Neighbors = append(node1.Neighbors, node1)
	node2.Neighbors = append(node1.Neighbors, node3)
	node3.Neighbors = append(node1.Neighbors, node2)
	node3.Neighbors = append(node1.Neighbors, node4)
	node4.Neighbors = append(node1.Neighbors, node1)
	node4.Neighbors = append(node1.Neighbors, node3)
	result := cloneGraph(node1)
	print(result)
}
