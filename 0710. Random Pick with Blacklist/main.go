package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	BlackList []int
	BlackMap  map[int]int
	N         int
}

func Constructor(n int, blacklist []int) Solution {
	var s Solution
	s.N = n
	//sort.Ints(blacklist)
	s.BlackMap = make(map[int]int)
	for _, b := range blacklist {
		s.BlackMap[b] = -1
	}
	last := n - 1
	for _, v := range blacklist{
		if v >= n - len(blacklist){
			continue
		}
		_, ok := s.BlackMap[last]
		for ok {
			last--
			_, ok = s.BlackMap[last]
		}
		s.BlackMap[v] = last
		last--
	}
	return s
}

func (this *Solution) Pick() int {
	whiteLen := this.N - len(this.BlackMap)
	pos := rand.Intn(whiteLen)
	if v, ok := this.BlackMap[pos]; ok {
		return v
	} else {
		return pos
	}
}

func main() {
	obj := Constructor(10, []int{0,1,4,7,9})
	fmt.Println(obj)
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
	fmt.Println(obj.Pick())
}
