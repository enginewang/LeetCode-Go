package main

import "fmt"

func modifyString(str string) string {
	s := []byte(str)
	for i, c := range s {
		empty := 'a'
		if c == '?' {
			if i == 0 {
				if len(s) > 2 && s[i+1] == 'a' {
					empty = 'b'
				}
			} else {
				if s[i-1] == 'a' || (i < len(s)-1 && s[i+1] == 'a') {
					empty += 1
					if s[i-1] == 'b' || (i < len(s)-1 && s[i+1] == 'b') {
						empty += 1
					}
				}
			}
			s[i] = byte(empty)
		}
	}
	return string(s)
}

func main() {
	fmt.Println(modifyString("?adf"))
}
