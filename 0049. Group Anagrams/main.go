package main

func groupAnagrams(strs []string) [][]string {
	vocDict := make(map[[26]int] []string)
	for _, str := range strs {
		d := [26]int{}
		for _, r := range []rune(str) {
			d[r-'a']++
		}
		vocDict[d] = append(vocDict[d], str)
	}
	var result [][]string
	for _, v := range vocDict{
		result = append(result, v)
	}
	return result
}

func main() {
	groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
}
