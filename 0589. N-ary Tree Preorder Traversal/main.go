package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	curLevel := []Node{*root}
	for len(curLevel) > 0 {
		var curVal []int
		for _, i := range curLevel {
			curVal = append(curVal, i.Val)
		}
		result = append(result, curVal)
		var nextLevel []Node
		for _, i := range curLevel {
			for _, j := range i.Children{
				nextLevel = append(nextLevel, *j)
			}
		}
		curLevel = nextLevel
	}
	return result
}

func main() {
	node5 := Node{Val: 5, Children: nil}
	node6 := Node{Val: 6, Children: nil}
	node2 := Node{Val: 2, Children: nil}
	node4 := Node{Val: 4, Children: nil}
	node3 := Node{Val: 3, Children: []*Node{&node5, &node6}}
	node1 := Node{Val: 1, Children: []*Node{&node3, &node2, &node4}}
	result := levelOrder(&node1)
	fmt.Println(result)
}
