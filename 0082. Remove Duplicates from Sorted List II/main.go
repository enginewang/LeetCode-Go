package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}


func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	DummyNode := &ListNode{Val: -1, Next: head}
	left := DummyNode
	right := head
	for right.Next != nil {
		findSame := false
		for right.Next.Val == right.Val {
			findSame = true
			if right.Next.Next != nil {
				right = right.Next
			} else {
				left.Next = nil
				return DummyNode.Next
			}

		}
		if findSame {
			left.Next = right.Next
			right = right.Next
		} else {
			right = right.Next
			left = left.Next
		}
	}
	return DummyNode.Next
}

func main() {
	node8 := ListNode{5, nil}
	node7 := ListNode{5, &node8}
	node6 := ListNode{4, &node7}
	//node5 := ListNode{3, &node6}
	//node4 := ListNode{3, &node5}
	//node3 := ListNode{2, &node4}
	//node2 := ListNode{1, &node3}
	//node1 := ListNode{1, &node2}
	r := deleteDuplicates(&node6)
	fmt.Println(r)
}
