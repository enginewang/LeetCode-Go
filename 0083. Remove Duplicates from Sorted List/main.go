package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	for cur != nil && cur.Next != nil {
		next := cur.Next
		for next != nil && next.Val == cur.Val {
			next = next.Next
		}
		cur.Next = next
		cur = next
	}
	return head
}

func main() {
	node3 := ListNode{3, nil}
	node5 := ListNode{3, &node3}
	node4 := ListNode{2, &node5}
	node2 := ListNode{1, &node4}
	node1 := ListNode{1, &node2}
	deleteDuplicates(&node1)
}