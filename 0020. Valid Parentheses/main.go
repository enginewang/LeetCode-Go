package main

import "fmt"

type Stack struct {
	S []rune
	Top rune
}

func (s *Stack) Push(r rune)  {
	s.S = append(s.S, r)
	s.Top = r
}

func (s *Stack) Pop() rune {
	if len(s.S) == 0{
		return 0
	} else {
		top := s.Top
		s.S = s.S[0:len(s.S)-1]
		if len(s.S) > 0{
			s.Top = s.S[len(s.S)-1]
		} else{
			s.Top = 0
		}
		return top
	}
}

func isValid(s string) bool {
	var st Stack
	for r := range s{
		switch s[r] {
		case '(':
			st.Push('(')
		case ')':
			pop := st.Pop()
			if pop != '(' {
				return false
			}
		case '[':
			st.Push('[')
		case ']':
			pop := st.Pop()
			if pop != '[' {
				return false
			}
		case '{':
			st.Push('{')
		case '}':
			pop := st.Pop()
			if pop != '{' {
				return false
			}
		}
	}
	if st.Top != 0{
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("([)]"))
}
