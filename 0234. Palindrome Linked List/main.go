package main

// 要想达到O(1)的空间复杂度，方式是翻转后半个链表，然后就可以比较了

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	mid, end := head, head
	for end.Next != nil && end.Next.Next != nil{
		mid = mid.Next
		end = end.Next.Next
	}
	if end.Next != nil {
		mid = mid.Next
	}
	if mid.Next == nil{
		if mid.Val == head.Val{
			return true
		} else {
			return false
		}
	}
	back := mid
	mid = mid.Next
	back.Next = nil
	for mid.Next != nil {
		next := mid.Next
		mid.Next = back
		back = mid
		mid = next
	}
	mid.Next = back
	// mid和head比较
	for mid != nil {
		if head.Val != mid.Val{
			return false
		}
		head = head.Next
		mid = mid.Next
	}
	return true
}
func main() {
	node6 := ListNode{1, nil}
	node5 := ListNode{2, &node6}
	node4 := ListNode{3, &node5}
	node3 := ListNode{3, &node4}
	node2 := ListNode{2, &node3}
	node1 := ListNode{1, &node2}
	isPalindrome(&node1)
}
