package main

import "fmt"

// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	sList := []byte(s)
	left, right := 0, 0
	m := make(map[byte]bool)
	maxLength := 0
	for ; right < len(sList); right++ {
		// 如果在里面，也就是之前出现了，left一直右移
		if m[sList[right]] {
			for sList[left] != sList[right] {
				delete(m, sList[left])
				left++
			}
			left++
		} else {
			// 否则就直接进
			m[sList[right]] = true
			if right-left+1 > maxLength {
				maxLength = right - left + 1
			}
		}
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
