package leetcode

// plusOne
// 加一
func plusOne(digits []int) []int {
	num := len(digits)
	if num < 1 {
		return digits
	}

	for i := num - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	rs := make([]int, num+1)
	rs[0] = 1
	return rs
}
