package week03

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1.0 / pow(x, -n)
	}
	return pow(x, n)
}
func pow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	y := pow(x, n/2)
	if n%2 == 1 {
		return y * y * x
	}
	return y * y
}
