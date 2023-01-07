package main

// 很容易想到用栈实现
func removeDuplicates(s string) string {
	var st []byte
	for _, b := range []byte(s) {
		if len(st) > 0 && st[len(st)-1] == b {
			st = st[:len(st)-1]
		} else {
			st = append(st, b)
		}
	}
	return string(st)
}

func main() {

}
