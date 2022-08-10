package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	s1, s2 := 0, 0
	for s1 < len(g) && s2 < len(s) {
		if s[s2] >= g[s1] {
			s1 += 1
			s2 += 1
		} else {
			s2 += 1
		}
	}
	return s1
}

func main() {

}
