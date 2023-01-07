package main

/*
给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。
如果可以，返回 true ；否则返回 false

类似于有效的字母异位词的改版，不过这里magazine可以更多。

构造hashMap[k]v，k是magazine的字符，v是出现的个数。先遍历magazine完成这个map，然后遍历ransonNote，每次-1，如果小于0则返回false，否则true。
*/

//func canConstruct(ransomNote string, magazine string) bool {
//	m := make(map[rune]int)
//	for _, i := range magazine {
//		m[i] += 1
//	}
//	for _, j := range ransomNote {
//		m[j] -= 1
//		if m[j] < 0 {
//			return false
//		}
//	}
//	return true
//}

// 这里字符串由a-z构成，所以可以用一个[26]int的数组，这样更快而且节约空间。

func canConstruct(ransomNote string, magazine string) bool {
	m := make([]int, 26)
	for _, i := range magazine {
		m[i-'a'] += 1
	}
	for _, j := range ransomNote {
		if m[j-'a'] == 0 {
			return false
		}
		m[j-'a'] -= 1
	}
	return true
}

func main() {

}
