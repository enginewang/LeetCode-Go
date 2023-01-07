package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	//for root != nil && root.Val > k {
	//	root = root.Left
	//}
	if root == nil {
		return false
	}
	m := make(map[int]struct{})
	result := false
	find(root, k, &m, &result)
	return result
}

func find(root *TreeNode, k int, m *map[int]struct{}, result *bool) {
	if root == nil || *result == true {
		return
	}
	if _, ok := (*m)[root.Val]; ok {
		*result = true
		return
	}
	(*m)[k-root.Val] = struct{}{}
	left, right := root.Left, root.Right
	find(left, k, m, result)
	find(right, k, m, result)
}

func main() {

}
