package common

import "math"

func isPalindrome(x int) bool {
	// 负数肯定不是回文，因为负号
	if x < 0 {
		return false
	}
	// 记录下原始数字
	originX := x

	// 反转数字
	rev := 0
	for x != 0 {
		// 每次取余数 用于反转相加
		pop := x % 10
		// 递归除以10 进行向右偏移
		x = x / 10
		rev = rev*10 + pop
	}

	// 原始数字和反转数字对比
	if originX == rev {
		return true
	}
	return false
}

// 整数反转
func reverseNumber(x int) int {
	y := 0
	for x != 0 {
		y = y*10 + x%10
		if !(-(1<<31) <= y && y <= (1<<31)-1) {
			return 0
		}
		x /= 10
	}
	return y
}

// 整数反转
func reverseNumberV2(x int) int {
	ret := 0
	for x != 0 {
		pop := x % 10
		x = x / 10
		ret = ret*10 + pop
		if ret < math.MinInt32 || ret > math.MaxInt32 {
			return 0
		}
	}
	return ret
}
