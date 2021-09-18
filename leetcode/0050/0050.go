package leetcode

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

// 快速幂——迭代
// 时间复杂度：O(\log n)O(logn)，即为对 nn 进行二进制拆分的时间复杂度。
// 空间复杂度：O(1)O(1)。
func quickMul(x float64, n int) float64 {
	rs := 1.0

	// 初始值为x
	data := x

	// 0000
	// 8421
	// 0101

	// 在对n进行二进制拆分的同时计算答案
	for n > 0 {

		// 判断n是否能整除2
		// n为奇数时首次和最后一次满足条件
		// n为偶数时最后一次满足条件
		if n%2 == 1 {
			rs *= data
		}

		// 将贡献不断地平方
		data *= data

		// 舍弃N二进制表示的最低位，这样我们每次只要判断最低位即可
		n /= 2
	}
	return rs
}

// 快速幂——递归
func quickMulRecursion(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickMulRecursion(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}
