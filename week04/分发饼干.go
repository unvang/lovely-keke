package week04

import "sort"
//1.用大饼干优先满足胃口大的，并统计满足小孩数量；排序后从后往前遍历
//2.或者小饼干先喂饱小胃口
//所以要先排序
//1，3，5，9  饼干
//1，2，7，10 小孩
func findContentChildren(greed, size []int) (ans int) {
	sort.Ints(greed)
	sort.Ints(size)
	n, m := len(greed), len(size)
	for i, j := 0, 0; i < n && j < m; i++ {
		for j < m && greed[i] > size[j] {
			j++
		}
		if j < m {
			ans++
			j++
		}
	}
	return
}
