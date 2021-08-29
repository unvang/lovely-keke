package week04


//只有暴力解的思路
//看了题解，分析过程自己想不出来，动态规划和贪心两种解法，需要多做几遍
//多看几遍吧
func maxProfit(prices []int) (ans int) {
	for i := 1; i < len(prices); i++ {
		ans += max(0, prices[i]-prices[i-1])
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
