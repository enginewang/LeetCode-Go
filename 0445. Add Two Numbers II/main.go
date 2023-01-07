package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var s1, s2 []int
	p1, p2 := l1, l2
	for p1 != nil {
		s1 = append(s1, p1.Val)
		p1 = p1.Next
	}
	for p2 != nil {
		s2 = append(s2, p2.Val)
		p2 = p2.Next
	}
	if len(s1) < len(s2) {
		s1, s2 = s2, s1
	}
	diff := len(s1) - len(s2)
	for diff > 0 {
		s2 = append([]int{0}, s2...)
		diff -= 1
	}
	dummyNode := &ListNode{Next: nil}
	cur := dummyNode.Next
	nextAdd, add := 0, 0
	for len(s1) > 0 {
		add = s1[len(s1)-1] + s2[len(s2)-1]
		s1, s2 = s1[:len(s1)-1], s2[:len(s2)-1]
		node := &ListNode{Val: (add + nextAdd) % 10}
		if add+nextAdd >= 10 {
			nextAdd = 1
		} else {
			nextAdd = 0
		}
		node.Next = cur
		dummyNode.Next = node
		cur = node
	}
	if nextAdd == 1 {
		node := &ListNode{Val: 1, Next: dummyNode.Next}
		dummyNode.Next = node
	}
	return dummyNode.Next
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

func main() {
	l1 := ListToLinkedList([]int{1})
	l2 := ListToLinkedList([]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9})
	addTwoNumbers(l1, l2)
}
