package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if p == nil || q == nil || p.Val != q.Val {
        return false
    }
    if isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) {
        return true
    } else {
        return false
    }
}

func main() {

}
