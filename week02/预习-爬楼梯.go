package week02

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	n1, n2 := 1, 2
	res := 0
	for i := 3; i <= n; i++ {
		res = n1 + n2
		n1 = n2
		n2 = res
	}
	return res
}

//超时解法，如何优化？？？
func climbStairs(n int) int {
	return cacl(n)

}
func cacl(n int) int {
	if n <= 2 {
		return n
	}
	return cacl(n-2) + cacl(n-1)
}
