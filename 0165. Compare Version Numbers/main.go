package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	n := 0
	if len(v1) > len(v2) {
		n = len(v1)
	} else {
		n = len(v2)
	}
	for i := 0; i < n; i++ {
		i1, i2 := 0, 0
		if i < len(v1) {
			i1, _ = strconv.Atoi(v1[i])
		}
		if i < len(v2) {
			i2, _ = strconv.Atoi(v2[i])
		}
		if i1 > i2 {
			return 1
		} else if i1 < i2 {
			return -1
		}
	}
	return 0
}

func main() {
	fmt.Println(compareVersion("1.0", "1.0.0.1"))
}
