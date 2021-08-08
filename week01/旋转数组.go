package week01

//https://leetcode-cn.com/problems/rotate-array/solution/jiong-yao-fei-shi-duan-hua-chang-shuo-tu-v1rw/

//solution 1 空间复杂度O(1)
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}
func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

//solution 2
func rotate2(nums []int, k int) {
	n := len(nums)
	k = k % n
	newNum := make([]int, n)
	j := 0
	for i := n - k; i < n; i++ {
		newNum[j] = nums[i]
		j++
	}
	for i := 0; i < n-k; i++ {
		newNum[j] = nums[i]
		j++
	}
	copy(nums, newNum)
}
