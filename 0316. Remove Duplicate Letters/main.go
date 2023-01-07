package main

// 单调栈
func removeDuplicateLetters(s string) string {
	var m [26]int
	var inStack [26]bool
	for _, b := range []byte(s) {
		m[b-'a'] += 1
	}
	// 单调栈
	var stack []byte
	for _, char := range []byte(s) {
		if !inStack[char-'a'] {
			// 发生这些情况，出栈
			for len(stack) > 0 && stack[len(stack)-1] >= char && m[stack[len(stack)-1]-'a'] > 0 {
				inStack[stack[len(stack)-1]-'a'] = false
				stack = stack[:len(stack)-1]
				//canRemove = true
			}
			stack = append(stack, char)
			inStack[char-'a'] = true
		}
		m[char-'a'] -= 1
	}
	return string(stack)
}

func main() {
	removeDuplicateLetters("cbacdcbc")
}
