package main

import "fmt"

type Trie struct {
	isEnd bool
	Child [26]*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	cur := this
	for _, b := range []byte(word) {
		b -= 'a'
		if cur.Child[b] == nil {
			cur.Child[b] = &Trie{isEnd: false}
		}
		cur = cur.Child[b]
	}
	cur.isEnd = true
}

func (this *Trie) Search(word string) bool {
	cur := this
	for _, b := range []byte(word) {
		b -= 'a'
		if cur.Child[b] == nil {
			return false
		}
		cur = cur.Child[b]
	}
	return cur.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, b := range []byte(prefix) {
		b -= 'a'
		if cur.Child[b] == nil {
			return false
		}
		cur = cur.Child[b]
	}
	return true
}

func main() {
	t := Constructor()
	t.Insert("apple")
	fmt.Print(t)
}
