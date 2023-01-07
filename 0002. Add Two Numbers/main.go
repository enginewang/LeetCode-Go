package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	head := &ListNode{}
	cur := head
	for !(p1.Next == nil && p2.Next == nil) {
		if p1.Next == nil {
			node := ListNode{Val: 0}
			p1.Next = &node
		}
		if p2.Next == nil {
			node := ListNode{Val: 0}
			p2.Next = &node
		}
		p1, p2 = p1.Next, p2.Next
	}
	p1, p2, add := l1, l2, 0
	for p1 != nil && p2 != nil {
		node := &ListNode{Val: p1.Val + p2.Val + add}
		cur.Next = node
		cur = cur.Next
		if cur.Val >= 10 {
			add = cur.Val / 10
			cur.Val = cur.Val % 10
		} else {
			add = 0
		}
		p1, p2 = p1.Next, p2.Next
	}
	if add > 0 {
		cur.Next = &ListNode{Val: add, Next: nil}
	}
	return head.Next
}

func main() {

}
