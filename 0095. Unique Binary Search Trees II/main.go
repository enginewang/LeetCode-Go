package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type StartEnd struct {
	start int
	end int
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	memo := make(map[StartEnd] []*TreeNode)
	return helper(StartEnd{start: 1, end: n}, &memo)
}

func helper(s StartEnd, memo *map[StartEnd] []*TreeNode) []*TreeNode {
	if s.start > s.end {
		return []*TreeNode{nil}
	}
	allTrees := []*TreeNode{}
	if (*memo)[s] != nil {
		return (*memo)[s]
	}
	for i := s.start; i <= s.end; i++ {
		leftTrees := helper(StartEnd{s.start, i - 1}, memo)
		rightTrees := helper(StartEnd{i + 1, s.end}, memo)
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				currTree := &TreeNode{i, nil, nil}
				currTree.Left = left
				currTree.Right = right
				allTrees = append(allTrees, currTree)
			}
		}
		(*memo)[s] = allTrees
	}
	return allTrees
}

func main() {

}
