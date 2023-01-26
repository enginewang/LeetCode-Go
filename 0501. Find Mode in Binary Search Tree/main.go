package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	var prev *TreeNode
	result := make([]int, 0)
	count, max := 1, 1
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		if prev != nil && prev.Val == node.Val {
			count++
		} else {
			count = 1
		}
		if count >= max {
			if count > max && len(result) > 0 {
				result = []int{node.Val}
			} else {
				result = append(result, node.Val)
			}
			max = count
		}
		prev = node
		traverse(node.Right)
	}
	traverse(root)
	return result
}

func main() {

}
