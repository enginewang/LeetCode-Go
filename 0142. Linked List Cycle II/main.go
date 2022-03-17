package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//func detectCycle(head *ListNode) *ListNode {
//	m := make(map[*ListNode] bool)
//	count := 0
//	for head != nil {
//		if _, ok := m[head]; ok{
//			return head
//		}
//		m[head] = true
//		head = head.Next
//		count++
//	}
//	return nil
//}

//
//func detectCycle(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//	slow, fast := head, head
//	for fast.Next != nil && fast.Next.Next != nil {
//		fast = fast.Next.Next
//		slow = slow.Next
//		if slow == fast {
//			slow = head
//			for slow != fast {
//				slow = slow.Next
//				fast = fast.Next
//			}
//			return slow
//		}
//	}
//	return nil
//}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow{
			for slow != head{
				slow = slow.Next
				head = head.Next
			}
			return slow
		}
	}
	return nil
}

func main() {
	node4 := ListNode{4, nil}
	node3 := ListNode{3, &node4}
	node2 := ListNode{2, &node3}
	node1 := ListNode{1, &node2}
	node4.Next = &node2
	detectCycle(&node1)
}
