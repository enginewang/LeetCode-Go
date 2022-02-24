package main

import (
	"fmt"
)

type Queue struct {
	Q []string
}

func (q *Queue) Len() int {
	return len(q.Q)
}

func (q *Queue) Push(s string) {
	q.Q = append(q.Q, s)
}

func (q *Queue) Pop() string {
	first := q.Q[0]
	q.Q = q.Q[1:]
	return first
}

func getNeighbor(s string) []string {
	var result []string
	c := []byte(s)
	tmp := make([]byte, len(s))
	copy(tmp, c)
	for i := 0; i < 4; i ++ {
		if c[i] == '0'{
			c[i] = '9'
		} else {
			c[i] = c[i] - 1
		}
		result = append(result, string(c[:]))
		copy(c, tmp)
		if c[i] == '9' {
			c[i] = '0'
		} else {
			c[i] = c[i] + 1
		}
		result = append(result, string(c[:]))
		copy(c, tmp)
	}
	return result
}

func openLock(deadends []string, target string) int {
	var q Queue
	result := 0
	deadMap := make(map[string]bool)
	for _, v := range deadends { // 记录所有“死亡点”
		deadMap[v] = true
	}
	// 参观过的号码
	visited := make(map[string] struct{})
	first := "0000"
	q.Push(first)
	//visited[first] = struct{}{}
	for q.Len() > 0 {
		size := q.Len()
		for i := 0; i < size; i++{
			cur := q.Pop()
			if _, ok := visited[cur]; ok {
				continue
			} else {
				visited[cur] = struct{}{}
			}
			if cur == target {
				return result
			}
			if _, ok := deadMap[cur]; ok{
				continue
			}
			for _, neighbor := range getNeighbor(cur){
				q.Push(neighbor)
			}
		}
		result++
	}
	return -1
}

func main() {
	fmt.Println(openLock([]string{"8887","8889","8878","8898","8788","8988","7888","9888"}, "8888"))
}
