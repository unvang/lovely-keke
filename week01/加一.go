package week01

func plusOne(digits []int) []int {
	n := len(digits)
	if n == 0 {
		return digits
	}
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] = digits[i] + 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	newDigits := make([]int, len(digits)+1)
	newDigits[0] = 1
	return newDigits
}
