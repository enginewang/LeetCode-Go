package main

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

type nl []*ListNode

func (l nl) Len() int {
	return len(l)
}

func (l nl) Less(i int, j int) bool {
	if l[i].Val-l[j].Val < 0 {
		return true
	} else {
		return false
	}
}

func (l nl) Swap(i int, j int) {
	l[i], l[j] = l[j], l[i]
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var nodeList nl
	for cur := head; cur != nil; cur = cur.Next {
		nodeList = append(nodeList, cur)
	}
	sort.Sort(nodeList)
	for i := 0; i < len(nodeList)-1; i++ {
		nodeList[i].Next = nodeList[i+1]
	}
	nodeList[len(nodeList)-1].Next = nil
	return nodeList[0]
}

func main() {

}
