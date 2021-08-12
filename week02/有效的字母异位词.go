package week02

import (
	"sort"
	"strings"
)

//1	排序后比较
func isAnagram(s string, t string) bool {
	if SortString(s) != SortString(t) {
		return false
	}
	return true
}
func SortString(src string) string {
	s := strings.Split(src, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

//2 map
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counter := make(map[byte]int8)
	for i := 0; i < len(s); i++ {
		counter[s[i]]++
		counter[t[i]]--
	}
	for _, v := range counter {
		if v != 0 {
			return false
		}
	}

	return true
}

//3 因为s,t只包含小写字母,'a' asc2 是 97
func isAnagram3(s string, t string) bool {
	alphbet := make([]int, 26)
	for i, v := range s {
		alphbet[v-'a']++
		alphbet[t[i]-'a']--
	}
	for _, v := range alphbet {
		if v != 0 {
			return false
		}
	}
	return true
}
