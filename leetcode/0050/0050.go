package leetcode

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	rs := 1.0

	// 初始值为x
	data := x

	// 在对n进行二进制拆分的同时计算答案
	for n > 0 {

		// 如果n二进制表示的最低位为1，那么需要计入贡献
		if n%2 == 1 {
			rs *= data
		}
		// 将贡献不断地平方
		data *= data
		// 舍弃 N 二进制表示的最低位，这样我们每次只要判断最低位即可
		n /= 2
	}
	return rs
}
