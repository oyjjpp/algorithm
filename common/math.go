// 数据 & 位运算
package common

import "strconv"

// singleNumber
// 只出现一次的数字
// 思路：使用按位异或云算付"^"，
// 参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1
func singleNumber(nums []int) int {
	var rs int
	for _, v := range nums {
		rs = rs ^ v
	}
	return rs
}

// maxPoints
// 直线上最多的点数
// 给定一个二维平面，平面上有 n 个点，求最多有多少个点在同一条直线上
func maxPoints(points [][]int) int {
	return 0
}

// fractionToDecimal
// 分数到小数
// 1、符号的问题，结果是正号还是负号
// 2、整除还是有余数
// 3、余数够不够除数，以至于是否需要补零
// 4、小数点后第几位开始循环
func fractionToDecimal(numerator int, denominator int) string {
	s := ""
	n, d := numerator, denominator

	// 判断是正数还是负数
	if (n > 0 && d < 0) || (n < 0 && d > 0) {
		s += "-"
	}

	// 如果分子为负数 则进行转换
	if n < 0 {
		n *= -1
	}
	// 如果分母为负数 则进行转换
	if d < 0 {
		d *= -1
	}

	// 相除取正数部分并转换成string
	s += strconv.Itoa(n / d)
	// 如果可以正数则直接返回
	if n%d == 0 {
		return s
	}

	// 处理不能整除情况
	// 3/9
	// 0.3333333333333333
	s += "."

	// 余数小于分母则进行补零 例如 3/90
	n = n % d
	n *= 10
	for n < d {
		s += "0"
		n *= 10
	}

	// 保存余数 用于校验是否存在死循环
	modMap := make(map[int]int)
	// 保存余数
	var buf []int

	buf = append(buf, n/d)
	modMap[n%d] = 1
	n %= d

	isLoop := false
	start := -1
	for n%d != 0 {
		n *= 10
		buf = append(buf, n/d)
		idx, ok := modMap[n%d]
		if !ok {
			modMap[n%d] = len(buf)
		} else {
			// 死循环
			isLoop = true
			// 标记循环开始位置
			start = idx
			break
		}
		n %= d
	}

	// 将小数部分拼接
	for i := 0; i < len(buf); i++ {
		if i == start && isLoop {
			s += "("
		}
		s += strconv.Itoa(buf[i])
	}

	if isLoop {
		s += ")"
	}
	return s
}

// trailingZeroes
// 阶乘后的零
// 给定一个整数 n，返回 n! 结果尾数中零的数量
func trailingZeroes(n int) int {
	count := 0
	for n > 0 {
		count += n / 5
		n = n / 5
	}
	return count
}

// reverseBits
// 颠倒二进制位
// 颠倒给定的 32 位无符号整数的二进制位。
func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		// << 左移n位就是乘以2的n次方
		// & 每次取num的最后一位
		// | 通过按位或将结果赋予res
		// << 通过左移 res结果想高位移动
		res = (res << 1) | (num & 1)
		// >> 右移n位就是除以2的n次方
		num = num >> 1
	}
	return res
}

// hammingWeight
// 位1的个数
// 编写一个函数，输入是一个无符号整数，返回其二进制表达式中数字位数为 ‘1’ 的个数（也被称为汉明重量）
func hammingWeight(num uint32) int {
	res := 0
	for num > 0 {
		// 通过按位与“&” 判断最后一位是否为1
		if (num & 1) == 1 {
			res++
		}
		// 向右移动
		num = (num >> 1)
	}
	return res
}
