package utils

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListToLinkedList(l []int) *ListNode {
	dummyNode := &ListNode{}
	cur := dummyNode
	for i, _ := range l {
		node := &ListNode{Val: l[i]}
		cur.Next = node
		cur = cur.Next
	}
	cur.Next = nil
	return dummyNode.Next
}
