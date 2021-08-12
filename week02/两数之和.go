package week02

//map
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, v := range nums {
		if pos, ok := numMap[target-v]; ok {
			return []int{i, pos}
		} else {
			numMap[v] = i
		}
	}
	return nil
}

//暴力双循环
func twoSum2(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
