package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	Arr []int
	// 哈希表存储数组的value-position，从而能在O(1)的复杂度内找到元素的位置
	Map map[int] int
}


func Constructor() RandomizedSet {
	var rs RandomizedSet
	rs.Map = make(map[int]int)
	return rs
}


func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.Map[val]; ok {
		return false
	} else {
		this.Map[val] = len(this.Arr)
		this.Arr = append(this.Arr, val)
		return true
	}
}


func (this *RandomizedSet) Remove(val int) bool {
	if pos, ok := this.Map[val]; ok {
		this.Map[this.Arr[len(this.Arr)-1]] = pos
		this.Arr[pos], this.Arr[len(this.Arr)-1] = this.Arr[len(this.Arr)-1], this.Arr[pos]
		this.Arr = this.Arr[:len(this.Arr)-1]
		delete(this.Map, val)
		return true
	} else {
		return false
	}
}


func (this *RandomizedSet) GetRandom() int {
	randPos := rand.Intn(len(this.Arr))
	return this.Arr[randPos]
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Insert(1))
	fmt.Println(obj.Remove(2))
	fmt.Println(obj.Insert(2))
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.Remove(1))
	fmt.Println(obj.Insert(2))
	fmt.Println(obj.GetRandom())
	fmt.Println(obj)
}
