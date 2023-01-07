package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Key   int
	Value int
	Next  *Node
	Prev  *Node
}

type DoubleLinkedList struct {
	Head *Node
	Tail *Node
	Size  int
}


// 初始化双向链表
func InitDoubleLinkedList() DoubleLinkedList {
	var dll DoubleLinkedList
	dll.Head = &Node{Key: -1, Value: -1}
	dll.Tail = &Node{Key: -1, Value: -1}
	dll.Head.Next = dll.Tail
	dll.Tail.Prev = dll.Head
	dll.Size = 0
	return dll
}

// 插入最后一个
func (l *DoubleLinkedList) addLast(n *Node) {
	n.Next = l.Tail
	n.Prev = l.Tail.Prev
	l.Tail.Prev.Next = n
	l.Tail.Prev = n
	l.Size++
}

// 删除一个（一定存在）
func (l *DoubleLinkedList) removeNode(n *Node) {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
	l.Size--
}

// 删除第一个结点
func (l *DoubleLinkedList) removeFirst() (*Node, error) {
	if l.Head.Next == l.Tail{
		return &Node{}, errors.New("Empty")
	}
	firstNode := l.Head.Next
	firstNode.Prev.Next = firstNode.Next
	firstNode.Next.Prev = firstNode.Prev
	l.Size--
	return firstNode, nil
}

type LRUCache struct {
	HashMap map[int] *Node
	Cache DoubleLinkedList
	Capacity int
}

func Constructor(capacity int) LRUCache {
	var lruCache LRUCache
	lruCache.Capacity = capacity
	lruCache.HashMap = make(map[int] *Node)
	lruCache.Cache = InitDoubleLinkedList()
	return lruCache
}

// 提升某个已有的key为最近使用的
func (this *LRUCache) makeRecently(key int) {
	node := this.HashMap[key]
	this.Cache.removeNode(node)
	this.Cache.addLast(node)
}


// 增加一个新的Node到cache中
func (this *LRUCache) addRecently(key int, value int) {
	node := &Node{Key: key, Value: value}
	this.Cache.addLast(node)
	this.HashMap[key] = node
}


// 删除一个key对应的结点
func (this *LRUCache) deleteKey(key int) {
	node := this.HashMap[key]
	this.Cache.removeNode(node)
	delete(this.HashMap, key)
}


// 删除最后未使用的结点
func (this *LRUCache) removeLeastRecently() error{
	node, err := this.Cache.removeFirst()
	if err != nil {
		return err
	}
	//fmt.Printf("淘汰了%v: %v\n", node.Key, node.Value)
	delete(this.HashMap, node.Key)
	return nil
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.HashMap[key]; ok {
		// 数据提升为最近使用的
		this.makeRecently(key)
		return node.Value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.HashMap[key]; ok {
		node.Value = value
		this.makeRecently(key)
	} else {
		this.addRecently(key, value)
		if this.Cache.Size > this.Capacity {
			this.removeLeastRecently()
		}
	}
}

func main() {
	l := Constructor(2)
	l.Put(1,1)
	l.Put(2,2)
	fmt.Println(l.Get(1))
	l.Put(3,3)
	fmt.Println(l.Get(2))
	l.Put(4,4)
	fmt.Println(l.Get(1))
	fmt.Println(l.Get(3))
	fmt.Println(l.Get(4))
}
