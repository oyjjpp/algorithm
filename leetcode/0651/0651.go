package leetcode

// maxA
// 651 四键键盘
func maxA(n int) int {
	var dp func(n, sumA, cacheA int) int
	dp = func(n, sumA, cacheA int) int {
		// base case
		if n <= 0 {
			return sumA
		}
		// 状态转移
		a := dp(n-1, sumA+1, cacheA)
		b := dp(n-1, sumA+cacheA, cacheA)
		c := dp(n-2, sumA, sumA)
		sumA = max(max(a, b), c)
		return sumA
	}
	return dp(n, 0, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
