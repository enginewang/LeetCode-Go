package main


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var queue []TreeNode
	queue = append(queue, *root)
	var result [][]int
	var row []int
	for len(queue) != 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			head := queue[0]
			if head.Left != nil {
				queue = append(queue, *head.Left)
			}
			if head.Right != nil {
				queue = append(queue, *head.Right)
			}
			row = append(row, head.Val)
			queue = queue[1:]
		}
		result = append(result, row)
		row = []int{}
	}
	p, q := 0, len(result)-1
	for p < q{
		result[p], result[q] = result[q], result[p]
		p++
		q--
	}
	return result
}

func main() {

}
