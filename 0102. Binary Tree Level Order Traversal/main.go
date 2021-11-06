package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeWithDepth struct {
	Node  *TreeNode
	Depth int
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var queue []NodeWithDepth
	queue = append(queue, NodeWithDepth{Node: root, Depth: 0})
	var result [][]int
	var row []int
	for len(queue) != 0 {
		head := queue[0]
		if head.Node.Left != nil {
			queue = append(queue, NodeWithDepth{Node: head.Node.Left, Depth: head.Depth+1})
		}
		if head.Node.Right != nil {
			queue = append(queue, NodeWithDepth{Node: head.Node.Right, Depth: head.Depth+1})
		}
		row = append(row, head.Node.Val)
		// 下一个元素的depth在下一层，说明这个是本层的最后一个
		if (len(queue) > 1 && queue[1].Depth != head.Depth) || len(queue) == 1{
			result = append(result, row)
			row = []int{}
		}
		queue = queue[1:]
	}
	return result
}

func main() {
	Node1 := &TreeNode{Val: 3, Left: nil, Right: nil}
	Node2 := &TreeNode{Val: 9, Left: nil, Right: nil}
	Node3 := &TreeNode{Val: 20, Left: nil, Right: nil}
	Node4 := &TreeNode{Val: 15, Left: nil, Right: nil}
	Node5 := &TreeNode{Val: 7, Left: nil, Right: nil}
	Node1.Left = Node2
	Node1.Right = Node3
	Node3.Left = Node4
	Node3.Right = Node5
	fmt.Println(levelOrder(Node1))
}
