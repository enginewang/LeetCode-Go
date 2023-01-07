package main

import "fmt"

type Stack struct {
	S []int
}

func (s *Stack) Push(r int) {
	s.S = append(s.S, r)
}

func (s *Stack) Pop() bool {
	if len(s.S) == 0 {
		return false
	} else {
		s.S = s.S[0 : len(s.S)-1]
		return true
	}
}

func (s *Stack) Top() int {
	return s.S[len(s.S)-1]
}

func longestValidParentheses(s string) int {
	maxLength := 0
	curLength := 0
	var stack Stack
	stack.Push(-1)
	for i, b := range []byte(s) {
		if b == '(' {
			stack.Push(i)
		} else if len(stack.S) > 1{
			stack.Pop()
			curLength = i - stack.Top()
			if curLength > maxLength {
				maxLength = curLength
			}
		} else {
			stack.Pop()
			stack.Push(i)
		}
	}
	return maxLength
}

func main() {
	fmt.Println(longestValidParentheses("()(()"))
}

