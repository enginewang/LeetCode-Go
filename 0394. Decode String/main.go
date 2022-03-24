package main

import "fmt"

func decodeString(s string) string {
	arr := []byte(s)
	return string(decode(arr))
}

func decode(arr []byte) []byte {
	i := 0
	for arr[i]{

	}
}

func main() {
	fmt.Println(decodeString("3[a2[c]]"))
}
