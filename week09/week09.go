package week09

import "sort"

/*三角形最小路径和*/
//三角形最小路径和，最优代码
func minimumTotal(triangle [][]int) int {
	//从倒数第二行开始
	for r := len(triangle) - 2; r >= 0; r-- {
		for c := range triangle[r] {
			//复用triangle，将下一行的累加到当前行
			triangle[r][c] += min(triangle[r+1][c], triangle[r+1][c+1])
		}
	}
	return triangle[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//递归代码
func minimumTotal2(triangle [][]int) int {
	return dfs(triangle, 0, 0)
}
func dfs(triangle [][]int, i, j int) int {
	if i == len(triangle) {
		return 0
	}
	return min(dfs(triangle, i+1, j), dfs(triangle, i+1, j+1)) + triangle[i][j]
}

/*零钱兑换*/

func coinChange(coins []int, amount int) int {
	n := len(coins)
	//填充dp数组
	INF := amount + 1
	dp := make([]int, amount+1)
	for i, _ := range dp {
		dp[i] = INF
	}
	//填充
	dp[0] = 0
	//i是待合成的值，dp表示合成i需要多少个硬币
	for i := 1; i <= amount; i++ {
		for j := 0; j < n; j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] == INF {
		return -1
	}
	return dp[amount]
}

/*数组的相对排序*/

func relativeSortArray(arr1 []int, arr2 []int) []int {
	rank := map[int]int{}
	for i, v := range arr2 {
		rank[v] = i - len(arr2)
	}
	sort.Slice(arr1, func(i, j int) bool {
		x, y := arr1[i], arr1[j]
		if r, has := rank[x]; has {
			x = r
		}
		if r, has := rank[y]; has {
			y = r
		}
		return x < y
	})
	return arr1
}

/*有效的字母异位词*/

func isAnagram(s, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}

//map
func isAnagram2(s, t string) bool {
	var c1, c2 [26]int
	for _, ch := range s {
		c1[ch-'a']++
	}
	for _, ch := range t {
		c2[ch-'a']++
	}
	return c1 == c2
}
