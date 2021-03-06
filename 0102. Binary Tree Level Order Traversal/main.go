package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var queue []TreeNode
	queue = append(queue, *root)
	var result [][]int
	var row []int
	for len(queue) != 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			head := queue[0]
			if head.Left != nil {
				queue = append(queue, *head.Left)
			}
			if head.Right != nil {
				queue = append(queue, *head.Right)
			}
			row = append(row, head.Val)
			queue = queue[1:]
		}
		result = append(result, row)
		row = []int{}
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
