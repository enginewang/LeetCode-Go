package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	need, window := make(map[int] int), make(map[int] int)
	for _, tRune := range t {
		need[int(tRune)]++
	}
	left, right := 0, 0
	minWinSize := math.MaxInt32
	// valid表示满足need条件的字符个数，如果valid = len(need)，说明windows满足条件，完全覆盖了need
	valid := 0
	result := ""
	for right < len(s) {
		c := int(s[right])
		// 如果将移入的元素在need里面，window存储这个字符
		if need[c] > 0{
			window[c]++
			// 如果c字符的windows数量和need数量相同，说明这个字符满足了条件
			if window[c] == need[c]{
				valid += 1
			}
		}
		//fmt.Println(window)
		// 如果包含了所有的元素，就往右移动left
		for valid == len(need) && left <= right{
			// 如果出现更小的就更新minWinSize
			if (right - left + 1) < minWinSize{
				minWinSize = right - left + 1
				result = s[left:right+1]
			}
			// 只看原s有的元素，有的话就删掉
			if _, ok := need[int(s[left])]; ok {
				window[int(s[left])]--
				if window[int(s[left])] < need[int(s[left])]{
					valid -= 1
				}
			}
			left++
		}
		right += 1
	}
	return result
}


func main() {
	fmt.Println(minWindow("BADBC", "ABC"))
}
