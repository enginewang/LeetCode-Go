package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func preorderTraversal(root *TreeNode) []int {
//	var result []int
//	pTraversal(root, &result)
//	return result
//}
//
//func pTraversal(cur *TreeNode, result *[]int) {
//	if cur == nil {
//		return
//	}
//	*result = append(*result, cur.Val)
//	pTraversal(cur.Left, result)
//	pTraversal(cur.Right, result)
//}

// 迭代

func preorderTraversal(root *TreeNode) []int {
	var st []*TreeNode
	var result []int
	if root == nil {
		return result
	}
	st = append(st, root)
	for len(st) > 0 {
		cur := st[len(st)-1]
		result = append(result, cur.Val)
		st = st[:len(st)-1]
		if cur.Right != nil {
			st = append(st, cur.Right)
		}
		if cur.Left != nil {
			st = append(st, cur.Left)
		}
	}
	return result
}

func main() {

}
