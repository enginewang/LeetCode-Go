package main

// 设计链表

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

// 虚拟头结点
func Constructor() MyLinkedList {
	return MyLinkedList{Val: -1, Next: nil}
}

func (this *MyLinkedList) Get(index int) int {
	cur := this.Next
	for i := 0; i < index; i++ {
		if cur != nil {
			cur = cur.Next
		} else {
			return -1
		}
	}
	if cur == nil {
		return -1
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := MyLinkedList{Val: val, Next: this.Next}
	this.Next = &newNode
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := MyLinkedList{Val: val, Next: nil}
	cur := this
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &newNode
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		this.AddAtHead(val)
	} else {
		newNode := MyLinkedList{Val: val, Next: nil}
		cur := this
		for index > 0 {
			index--
			if cur == nil {
				return
			} else {
				cur = cur.Next
			}
		}
		if cur == nil {
			return
		}
		newNode.Next = cur.Next
		cur.Next = &newNode
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	cur := this
	for index > 0 {
		index--
		if cur == nil {
			return
		} else {
			cur = cur.Next
		}
	}
	if cur == nil || cur.Next == nil {
		return
	}
	cur.Next = cur.Next.Next
}

func main() {
	linkedList := Constructor()
	linkedList.AddAtHead(1)
	linkedList.Get(1)
}
