package main

import (
	"fmt"
)

func reverseWords(s string) string {
	// 通过双指针删除头部、中间、尾部空格
	slow, fast := 0, 0
	sList := []byte(s)
	for sList[fast] == ' ' {
		fast++
	}
	for ; fast < len(sList); fast++ {
		if fast > 0 && sList[fast] == ' ' && sList[fast-1] == ' ' {
			continue
		}
		sList[slow] = sList[fast]
		slow++
	}
	if sList[slow-1] == ' ' {
		slow--
	}
	sList = sList[:slow]
	// 整体颠倒
	left, right := 0, len(sList)-1
	for left < right {
		sList[left], sList[right] = sList[right], sList[left]
		left++
		right--
	}
	//fmt.Println(string(sList))
	// 单词颠倒
	start := 0
	for i, b := range sList {
		if b == ' ' || i == len(sList)-1 {
			left, right = start, i-1
			if i == len(sList)-1 {
				right = i
			}
			for left < right {
				sList[left], sList[right] = sList[right], sList[left]
				left++
				right--
			}
			start = i + 1
		}
	}
	return string(sList)
}

func main() {
	fmt.Println(reverseWords(" the sky     is blue    "))
}
