package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var Queue []*TreeNode
	var result [][]int
	Queue = append(Queue, root)
	level := 0
	for len(Queue) != 0 {
		var curLevel []int
		length := len(Queue)
		if level%2 == 0 {
			for i := 0; i < length; i++ {
				curLevel = append(curLevel, Queue[0].Val)
				if Queue[0].Left != nil {
					Queue = append(Queue, Queue[0].Left)
				}
				if Queue[0].Right != nil {
					Queue = append(Queue, Queue[0].Right)
				}
				Queue = Queue[1:]
			}
		} else {
			for i := length - 1; i >= 0; i-- {
				curLevel = append(curLevel, Queue[len(Queue)-1].Val)
				if Queue[len(Queue)-1].Right != nil {
					Queue = append([]*TreeNode{Queue[len(Queue)-1].Right}, Queue...)
				}
				if Queue[len(Queue)-1].Left != nil {
					Queue = append([]*TreeNode{Queue[len(Queue)-1].Left}, Queue...)
				}
				Queue = Queue[:len(Queue)-1]
			}
		}
		result = append(result, curLevel)
		level += 1
	}
	fmt.Println(result)
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
	zigzagLevelOrder(Node1)
}
