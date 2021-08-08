package week01

//solution 1
func moveZeroes(nums []int) {
	j := 0
	for i, v := range nums {
		if v != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
}

//solution 2
func moveZeroes2(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	j := 0
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for j < n {
		nums[j] = 0
		j++
	}
}
