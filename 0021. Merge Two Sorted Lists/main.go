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
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	head := &ListNode{}
	if l1.Val < l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}

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

	return head
}
func main() {
	node5 := ListNode{6, nil}
	node4 := ListNode{1, &node5}
	node3 := ListNode{-1, &node4}
	node2 := ListNode{-3, &node3}
	node1 := ListNode{-6, &node2}
	node0 := ListNode{-9, &node1}
	node7 := ListNode{5, nil}
	node8 := ListNode{-2, &node7}
	n := mergeTwoLists(&node0,&node8)
	fmt.Println(n)
}
