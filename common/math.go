// 数据 & 位运算
package common

import (
	"strconv"
	"strings"
)

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
// TODO
// 直线上最多的点数
// 给定一个二维平面，平面上有 n 个点，求最多有多少个点在同一条直线上
func maxPoints(points [][]int) int {
	return 0
}

// fractionToDecimal
// 分数到小数（string类型）
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

	// 如果是负数 通过（*-1） 操作转换为正数
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

	// 保存余数,即余数所在位置 用于校验是否存在死循环
	modMap := make(map[int]int)
	// 保存余数
	var buf []int

	buf = append(buf, n/d)
	modMap[n%d] = 1
	n %= d

	// 用于校验是否存在循环
	isLoop := false
	start := -1
	for n%d != 0 {
		n *= 10
		buf = append(buf, n/d)
		if idx, ok := modMap[n%d]; ok {
			// 死循环
			isLoop = true
			// 标记循环开始位置
			start = idx
			break
		} else {
			modMap[n%d] = len(buf)
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
// 思路
// 主要看当前数能被5整数的数据有多少
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
// 颠倒给定的32位无符号整数的二进制位。
// 思路
// 1、原数据向右侧推进
// 2、结果数据向右侧推进
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
// 思路
// 1、通过末尾按位与“&” 操作判断，如果是1 则计数加1
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

// 计数质数
// countPrimes[厄拉多塞筛法]
// 统计所有小于非负整数n的质数的数量
func countPrimes(n int) int {
	count := 0
	// 用来记录“已经找过的数的倍数”的
	signs := make([]bool, n)
	for i := 2; i < n; i++ {
		if signs[i] {
			continue
		}
		count++
		// 计算当前数据的倍数 都不可能为质数
		for j := 2 * i; j < n; j += i {
			signs[j] = true
		}
	}
	return count
}

// isPrime
// 是否为质数
// 思路
// 1、2和3是质数
// 2、能被2和3整除的都不是质数
// 3、
func isPrime(value int) bool {
	// 校验1/2/3等情况 1就不是质数也不是合数 2/3为质数
	if value <= 3 {
		return value >= 2
	}

	// 能够整除2/3都不是质数
	if value%2 == 0 || value%3 == 0 {
		return false
	}
	for i := 5; i*i <= value; i += 6 {
		if value%i == 0 || value%(i+2) == 0 {
			return false
		}
	}
	return true
}

// missingNumber
// 缺失数字
// 给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数
// 思路
// 只出现一次的数字，使用按位异或“^” 所有数据与索引进行按位异或后剩余的就是缺失的
func missingNumber(nums []int) int {
	// res赋值为n，因为循环时候索引会少一个n
	res := len(nums)
	for i := 0; i < len(nums); i++ {
		res = res ^ nums[i]
		res = res ^ i
	}
	return res
}

// 3的幂
// 给定一个整数，写一个函数来判断它是否是3的幂次方。
// 思路
// 1、循环除三 直到最后，如果除数为1，则代表符合，否则不符合
func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}

	// 循环除以3
	for n%3 == 0 {
		n = n / 3
	}
	return n == 1
}

// 将整数n转换为3进制字符串，判断是否第一位为1其他位为0，符合条件则为3的幂。
func isPowerOfThreeV2(n int) bool {
	if n < 1 {
		return false
	}
	s := strconv.FormatInt(int64(n), 3)
	return s[0:1] == "1" && strings.Count(s, "0") == len(s)-1
}
