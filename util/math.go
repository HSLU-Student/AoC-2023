package util

// greatest common divider of two numbers - implemented as eucalidian algorithm
func Gcd(a, b int64) int64 {
	//because everything is a divisor of zero
	if a == 0 {
		return b
	}

	for b != 0 {
		h := a % b
		a = b
		b = h
	}
	return a
}

// lowest common multiple of two numbers
func Lcm(a, b int64) int64 {
	return a * b / Gcd(a, b)
}

// abs of a int
func Abs(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}
