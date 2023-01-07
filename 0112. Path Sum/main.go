package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	hasFind := false
	if root == nil {
		return false
	}
	backtrack(root, root.Val, targetSum, &hasFind)
	return hasFind
}

func backtrack(node *TreeNode, curSum int, targetSum int, hasFind *bool) {
	if *hasFind == true {
		return
	}
	if node.Left == nil && node.Right == nil {
		if curSum == targetSum {
			*hasFind = true
			return
		} else {
			return
		}
	}
	if node.Left != nil {
		curSum += node.Left.Val
		backtrack(node.Left, curSum, targetSum, hasFind)
		curSum -= node.Left.Val
	}
	if node.Right != nil {
		curSum += node.Right.Val
		backtrack(node.Right, curSum, targetSum, hasFind)
		curSum -= node.Right.Val
	}
}

func main() {

}
