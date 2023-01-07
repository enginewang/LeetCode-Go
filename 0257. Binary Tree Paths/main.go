package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func binaryTreePaths(root *TreeNode) []string {
	var result []string
	backtrack(root, []string{}, &result)
	return result
}

func backtrack(cur *TreeNode, path []string, result *[]string)  {
	path = append(path, strconv.Itoa(cur.Val))
	if cur.Left == nil && cur.Right == nil {
		pathStr := strings.Join(path, "->")
		*result = append(*result, pathStr)
	} else {
		if cur.Left != nil {
			backtrack(cur.Left, path, result)
		}
		if cur.Right != nil {
			backtrack(cur.Right, path, result)
		}
	}
}

func main() {

}
