package week06

//这题有意思，看题解很容易理解，第一遍感觉记住怎么做就好了
//对比爬楼梯，fn=fn-1 + fn-2 ，dp[i][j]通过i,j两个坐标定位下一个值
//爬楼梯只需要 n定位，所以，此处dp记录二维数组，爬楼梯记录几个变量
//爬楼梯如果使用缓存的方式，也是缓存一个数组
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i, v1 := range text1 {
		for j, v2 := range text2 {
			if v1 == v2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[m][n]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
