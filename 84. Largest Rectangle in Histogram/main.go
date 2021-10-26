package main

import "fmt"

type Elem struct {
	id  int
	height int
}

type Stack struct {
	S []Elem
}

func (s *Stack) Push(r Elem)  {
	s.S = append(s.S, r)
}

func (s *Stack) Pop() {
	if len(s.S) == 1{
		fmt.Println("empty")
	} else {
		s.S = s.S[0:len(s.S)-1]
	}
}

func (s *Stack) Top() Elem{
	return s.S[len(s.S)-1]
}

// 左边界就是栈的下一个元素的下标
func largestRectangleArea(heights []int) int {
	// m是第i个的最佳长方形
	s := Stack{[]Elem{}}
	s.Push(Elem{id: -1})
	maxArea := 0
	for i, h := range heights {
		// 比当前栈顶高度大，进栈
		if h < s.Top().height{
			// 开始退栈，直到能放进去
			for h < s.Top().height{
				t := s.Top()
				area := t.height * (i-s.S[len(s.S)-2].id-1)
				if area > maxArea {
					maxArea = area
				}
				s.Pop()
			}
		}
		s.Push(Elem{id: i, height: h})
	}
	// 最后栈中的元素
	for len(s.S) > 1 {
		t := s.Top()
		area := t.height * (len(heights)-s.S[len(s.S)-2].id-1)
		if area > maxArea {
			maxArea = area
		}
		s.Pop()
	}

	return maxArea
}

func main() {
	height := []int{6, 7, 5, 2, 4, 5, 9, 3}
	fmt.Println(largestRectangleArea(height))
}
