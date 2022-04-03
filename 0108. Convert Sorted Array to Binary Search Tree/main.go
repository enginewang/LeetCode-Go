package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二分
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0], Left: nil, Right: nil}
	}
	left, right := 0, len(nums)-1
	mid := left + (left+right)/2
	root := &TreeNode{Val: nums[mid], Left: nil, Right: nil}
	if mid != left {
		root.Left = sortedArrayToBST(nums[left:mid])
	}
	if mid != right {
		root.Right = sortedArrayToBST(nums[mid+1:right+1])
	}
	return root
}

func main() {

}
