package leetcode

import "log"

// coinChange
// 零钱兑换
// stack overflow
func coinChange(coins []int, amount int) int {
	// 求最小值，所以初始化为正无穷
	maxInt := 1<<31 - 1

	var dp func(amount int) int
	dp = func(amount int) int {
		if amount == 0 {
			return 0
		}
		if amount < 0 {
			return -1
		}
		num := maxInt
		// [1, 2, 5]
		// 11
		for _, coin := range coins {
			// 子问题
			subProblem := dp(amount - coin)
			if subProblem == -1 {
				continue
			}
			num = min(num, 1+subProblem)
		}

		if num == maxInt {
			return -1
		} else {
			return num
		}
	}

	return dp(amount)
}

// min
// 获取两个整型最小值
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// coinChange
// 零钱兑换
// 通过备忘录解决重叠子问题
// 自顶向下
func coinChangeV2(coins []int, amount int) int {
	// 求最小值，所以初始化为正无穷
	maxInt := 1<<31 - 1
	data := map[int]int{}

	var dp func(amount int) int
	dp = func(amount int) int {
		// 从备忘录中获取
		if _, ok := data[amount]; ok {
			return data[amount]
		}
		if amount == 0 {
			return 0
		}
		if amount < 0 {
			return -1
		}
		num := maxInt
		// [1, 2, 5]
		// 11
		for _, coin := range coins {
			// 子问题
			subProblem := dp(amount - coin)
			if subProblem == -1 {
				continue
			}
			num = min(num, 1+subProblem)
		}

		// 存储到备忘录中
		if num == maxInt {
			data[amount] = -1
		} else {
			data[amount] = num
		}
		return data[amount]
	}

	return dp(amount)
}

// coinChange
// 零钱兑换
// 通过备DP table解决重叠子问题
// 超出时间限制
//
// [3] 2
// [2] 1
// [1,2,5] 11
// [2,5,10,1] 27
func coinChangeV3(coins []int, amount int) int {
	dp := map[int]int{}
	// base case
	dp[0] = 0
	number := amount + 1
	// 外层for循环在遍历状态的所有可能取值
	for i := 1; i < number; i++ {
		// 内层for缓存在求所有选择的最小值
		for _, coin := range coins {
			// 子问题无解
			subProblem := i - coin
			if subProblem < 0 {
				continue
			}
			// 设置默认值
			if _, ok := dp[i]; !ok {
				dp[i] = number
			}
			// 求最小的
			if _, ok := dp[subProblem]; ok {
				dp[i] = min(dp[i], 1+dp[subProblem])
			}
			log.Println(i, dp[i])
		}
	}

	if rs, ok := dp[amount]; !ok || rs == number {
		return -1
	} else {
		return dp[amount]
	}
}
