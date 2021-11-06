package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	tDict, sDict := make(map[int] int), make(map[int] int)
	for _, tRune := range t {
		tDict[int(tRune)]++
	}
	left := 0
	minWinSize := math.MaxInt32
	result := ""
	for right := 0; right < len(s); right++ {
		if right < len(s) && tDict[int(s[right])] > 0{
			sDict[int(s[right])]++
		}
		// 如果包含了所有的元素，就往右移动left
		for checkValid(sDict, tDict) && left <= right{
			if (right - left + 1) < minWinSize{
				minWinSize = right - left + 1
				result = s[left:right+1]
			}
			// 只看原dict有的元素
			if _, ok := tDict[int(s[left])]; ok {
				sDict[int(s[left])]--
			}
			left++
		}
	}
	return result
}

// 检查是否s包含t
func checkValid(sDict map[int] int, tDict map[int] int) bool {
	for k, v := range tDict {
		if sDict[k] < v{
			return false
		}
	}
	return true
}


func main() {
	fmt.Println(minWindow("BADBC", "ABC"))
}
