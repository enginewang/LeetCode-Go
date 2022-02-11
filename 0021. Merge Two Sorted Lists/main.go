package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	head := &ListNode{Val: -1}
	current := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			current = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			current = l2
			l2 = l2.Next
		}
	}
	if l1 == nil {
		current.Next = l2
	} else {
		current.Next = l1
	}
	return head.Next
}

func main() {
	node5 := ListNode{5, nil}
	node4 := ListNode{4, &node5}
	node3 := ListNode{3, &node4}
	node2 := ListNode{2, &node3}
	node1 := ListNode{1, &node2}
	node0 := ListNode{0, &node1}
	node7 := ListNode{9, nil}
	node8 := ListNode{-1, &node7}
	n := mergeTwoLists(&node0,&node8)
	fmt.Println(n)
}
