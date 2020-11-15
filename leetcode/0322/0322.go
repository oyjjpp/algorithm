package leetcode

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
func coinChangeV3(coins []int, amount int) int {
	// dp := make([]int, amount+1)
	dp := map[int]int{}
	// base case
	dp[0] = 0
	number := amount + 1
	// 外层for循环在遍历所有状态的所有取值
	for i := 1; i < number; i++ {
		// 内层for缓存在求所有选择的最小值
		for _, coin := range coins {
			// 子问题无解
			if (i - coin) < 0 {
				continue
			}
			if _, ok := dp[i-coin]; ok {
				dp[i] = min(number, 1+dp[i-coin])
			} else {
				dp[i] = number
			}
		}
	}

	if rs, ok := dp[amount]; !ok || rs == number {
		return -1
	} else {
		return dp[amount]
	}
}
