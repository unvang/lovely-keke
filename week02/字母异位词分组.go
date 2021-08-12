package week02

import "sort"

//排序
func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, v := range strs {
		s := []byte(v)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		m[sortedStr] = append(m[sortedStr], v)
	}
	rst := make([][]string, 0, len(m))
	for _, v := range m {
		rst = append(rst, v)
	}
	return rst
}

//计数
func groupAnagrams2(strs []string) [][]string {
	m := map[[26]int][]string{}
	for _, v := range strs {
		cnt := [26]int{}
		for _, c := range v {
			cnt[c-'a']++
		}
		m[cnt] = append(m[cnt], v)
	}
	rst := make([][]string, 0, len(m))
	for _, v := range m {
		rst = append(rst, v)
	}
	return rst
}
