package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left == nil {
			root = root.Right
		} else if root.Right == nil {
			root = root.Left
		} else {
			// 获得左子树最大的节点
			leftMax := root.Left
			for leftMax.Right != nil {
				leftMax = leftMax.Right
			}
			root.Left = deleteNode(root.Left, leftMax.Val)
			leftMax.Left = root.Left
			leftMax.Right = root.Right
			root = leftMax
		}
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func main() {

}
