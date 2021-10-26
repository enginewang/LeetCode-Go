package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyNode := &ListNode{Val: -1, Next: head}
	prev := dummyNode
	cur := dummyNode.Next
	next := cur.Next
	for next != nil {
		find := false
		for next != nil && cur.Val == next.Val {
			cur = cur.Next
			next = cur.Next
			find = true
		}
		if find{
			if next == nil{
				prev.Next = nil
				return dummyNode.Next
			}
			cur = cur.Next
			next = cur.Next
		} else{
			cur = cur.Next
			prev = prev.Next
			next = next.Next
		}
		prev.Next = cur
	}
	return dummyNode.Next
}

func main() {
	node7 := ListNode{5, nil}
	node6 := ListNode{4, &node7}
	node5 := ListNode{3, &node6}
	node4 := ListNode{3, &node5}
	node3 := ListNode{2, &node4}
	node2 := ListNode{1, &node3}
	node1 := ListNode{1, &node2}
	deleteDuplicates(&node1)
}