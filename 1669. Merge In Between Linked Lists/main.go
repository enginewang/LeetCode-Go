package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	// left, right分别是list1要删除的外面的两边
	p1, left, right := list1, list1, list1
	// 2,5
	for i := 0; i <= b; i++ {
		if i == a-1 {
			left = p1
		}
		p1 = p1.Next
	}
	right = p1
	left.Next = list2
	p2 := list2
	for p2.Next != nil {
		p2 = p2.Next
	}
	p2.Next = right
	return list1
}

func main() {

}
