package main

import (
	"container/heap"
	"fmt"
)

// 先实现最小堆
type minHeap [][2]int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num] += 1
	}
	h := &minHeap{}
	heap.Init(h)
	for key, value := range m {
		heap.Push(h, [2]int{key, value})
		// 如果堆比k大，就pop一个
		if h.Len() > k {
			heap.Pop(h)
		}
		//fmt.Println(h)
	}
	res := make([]int, k)
	// 反着放，因为每次Pop的都是最小的
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return res
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
