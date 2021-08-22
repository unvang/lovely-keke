package week03

func subsets(nums []int) (ans [][]int) {
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		//关键点，set上面放进去了，此处拿出来
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return
}
