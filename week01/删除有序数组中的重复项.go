package week01

//solution 1
func removeDuplicates1(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	slow := 0
	for fast, v := range nums {
		if nums[fast] != nums[slow] {
			nums[slow+1] = v
			slow++
		}
	}
	return slow + 1
}

//solution 2
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
