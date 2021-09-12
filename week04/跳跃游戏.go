package week04

//max（i+nums[i]） 是随机一个位置最远能到达的位置，记为rightMost
//遍历一遍，判断该值能否到达n-1即可
//注意点：i需要小于rightMost才判断，才可到达i
func canJump(nums []int) bool {
	n := len(nums)
	rightMost := 0
	for i, v := range nums {
		if i <= rightMost {
			rightMost = max(rightMost, i+v)
			if rightMost >= n-1 {
				return true
			}
		}
	}
	return false
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
