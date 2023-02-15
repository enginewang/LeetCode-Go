package main

import (
	"fmt"
	"strconv"
	"strings"
)

var result []string

func restoreIpAddresses(s string) []string {
	result = make([]string, 0)
	backtrack(0, []int{0}, s)
	return result
}

func backtrack(start int, temp []int, s string) {
	if len(temp) > 5 {
		return
	}
	if start == len(s) && len(temp) == 5 {
		validIpList := make([]string, 0)
		for k := 1; k < len(temp); k++ {
			validIpList = append(validIpList, s[temp[k-1]:temp[k]])
		}
		result = append(result, strings.Join(validIpList, "."))
	}

	for i := start + 1; i <= len(s); i++ {
		if isValidIp(s[start:i]) {
			temp = append(temp, i)
			backtrack(i, temp, s)
			temp = temp[:len(temp)-1]
		}
	}
}

func isValidIp(s string) bool {
	ip, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if ip > 255 {
		return false
	}
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	return true
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
}
