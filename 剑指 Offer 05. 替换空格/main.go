package main

import "fmt"

func replaceSpace(s string) string {
	l := []byte(s)
	newLen := 0
	for _, i := range l {
		newLen += 1
		if i == ' ' {
			newLen += 2
		}
	}
	//fmt.Println(newLen)
	newStr := make([]byte, newLen)
	j := len(l) - 1
	// 逆序
	for i := len(newStr) - 1; i >= 0; {
		if l[j] != ' ' {
			newStr[i] = l[j]
			j--
			i--
		} else {
			newStr[i] = '0'
			newStr[i-1] = '2'
			newStr[i-2] = '%'
			i = i - 3
			j--
		}
	}
	return string(newStr)
}

func main() {
	fmt.Println(replaceSpace("We are happy"))
}
