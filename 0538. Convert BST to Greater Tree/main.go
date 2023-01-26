package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var prev *TreeNode
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node.Right != nil {
			traverse(node.Right)
		}
		if prev != nil {
			node.Val += prev.Val
		}
		prev = node
		if node.Left != nil {
			traverse(node.Left)
		}
	}
	traverse(root)
	return root
}

func main() {

}
