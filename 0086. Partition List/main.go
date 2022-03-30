package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	dummyNode := &ListNode{Val: -1, Next: head}
	quick, slow := dummyNode, dummyNode
	for quick.Next != nil {
		for quick.Next != nil && quick.Next.Val >= x {
			quick = quick.Next
		}
		if quick.Next == nil {
			break
		}
		node := quick.Next
		if slow.Next != node {
			quick.Next = quick.Next.Next
			node.Next = slow.Next
			slow.Next = node
		} else {
			quick = quick.Next
		}
		slow = slow.Next
	}
	return dummyNode.Next
}

func main() {
	node6 := &ListNode{Val: 2, Next: nil}
	node5 := &ListNode{Val: 5, Next: node6}
	node4 := &ListNode{Val: 0, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 4, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}
	partition(node1, 2)
}
