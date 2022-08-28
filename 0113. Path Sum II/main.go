package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var result [][]int
	var curPath []int
	if root == nil {
		return result
	}
	curPath = append(curPath, root.Val)
	backtrack(root, curPath, root.Val, targetSum, &result)
	return result
}

func backtrack(node *TreeNode, curPath []int, curSum int, targetSum int, result *[][]int) {
	if node.Left == nil && node.Right == nil {
		if curSum == targetSum {
			tmp := make([]int, len(curPath))
			copy(tmp, curPath)
			*result = append(*result, tmp)
			return
		} else {
			return
		}
	}
	if node.Left != nil {
		curPath = append(curPath, node.Left.Val)
		curSum += node.Left.Val
		backtrack(node.Left, curPath, curSum, targetSum, result)
		curPath = curPath[:len(curPath)-1]
		curSum -= node.Left.Val
	}
	if node.Right != nil {
		curPath = append(curPath, node.Right.Val)
		curSum += node.Right.Val
		backtrack(node.Right, curPath, curSum, targetSum, result)
		curPath = curPath[:len(curPath)-1]
		curSum -= node.Right.Val
	}
}

func main() {

}
