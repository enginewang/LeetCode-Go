package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	var final []string
	l := strings.Split(path, "/")
	for _, elem := range l {
		fmt.Println(elem)
		if elem != "" && elem != "." && elem != ".."{
			final = append(final, elem)
		}
		if elem == ".." && len(final) > 0{
			final = final[:len(final)-1]
		}
	}
	return "/" + strings.Join(final, "/")
}

func main() {
	fmt.Println(simplifyPath("/a/./b/../../c/"))
}
