package week01

//solution 1
func twoSum(nums []int, target int) []int {
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

//solution 2
func twoSum2(nums []int, target int) []int {
	hashMap := make(map[int]int, 0)
	for i, v := range nums {
		if index, ok := hashMap[target-v]; ok {
			return []int{i, index}
		}
		hashMap[v] = i
	}
	return nil
}
