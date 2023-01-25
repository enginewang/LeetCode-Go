package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	pos := findMaxPos(nums)
	node := &TreeNode{Val: nums[pos]}
	leftChild := constructMaximumBinaryTree(nums[:pos])
	rightChild := constructMaximumBinaryTree(nums[pos+1:])
	node.Left = leftChild
	node.Right = rightChild
	return node
}

func findMaxPos(nums []int) int {
	maxPos, maxV := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxV {
			maxV = nums[i]
			maxPos = i
		}
	}
	return maxPos
}

func main() {
	node := constructMaximumBinaryTree([]int{1})
	fmt.Println(node)
}
