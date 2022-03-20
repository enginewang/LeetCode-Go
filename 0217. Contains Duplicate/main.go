package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		} else {
			m[num] = struct{}{}
		}
	}
	return false
}

func main() {

}
