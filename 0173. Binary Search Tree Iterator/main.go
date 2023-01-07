package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	cur  int
	list []int
}

func Constructor(root *TreeNode) BSTIterator {
	var list []int
	helper(root, &list)
	return BSTIterator{cur: -1, list: list}
}

func helper(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	helper(root.Left, list)
	*list = append(*list, root.Val)
	helper(root.Right, list)
}

func (this *BSTIterator) Next() int {
	this.cur += 1
	return this.list[this.cur]
}

func (this *BSTIterator) HasNext() bool {
	if this.cur < len(this.list)-1 {
		return true
	}else {
		return false
	}
}

func main() {

}
