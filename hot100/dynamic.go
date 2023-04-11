package hot100

import (
	"fmt"
)

// 动态规划
// 509. 斐波那契数
func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// 自顶向下【备忘录】
// 509. 斐波那契数
func fibV(n int) int {
	origin := map[int]int{}

	var dp func(data map[int]int, m int) int

	dp = func(data map[int]int, m int) int {
		if m == 0 || m == 1 {
			return m
		}

		if data[m] != 0 {
			return data[m]
		}
		data[m] = dp(data, m-1) + dp(data, m-2)
		return data[m]
	}
	return dp(origin, n)
}

// 自底向上
// 509. 斐波那契数
func fibV2(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 动态规划
// 322. 零钱兑换
// 自顶向下
func coinChange(coins []int, amount int) int {
	memo := make([]int, amount+1)
	// 将备忘录初始化为 -666，代表还未被计算
	for i := range memo {
		memo[i] = -666
	}
	maxNum := 1 << 20
	fmt.Println(maxNum)

	var dp func(coins []int, amount int) int
	dp = func(coins []int, amount int) int {
		if amount == 0 {
			return 0
		}
		if amount < 0 {
			return -1
		}

		if memo[amount] != -666 {
			return memo[amount]
		}

		res := maxNum
		for _, coin := range coins {
			subProblem := dp(coins, amount-coin)
			if subProblem == -1 {
				continue
			}
			res = min(res, 1+subProblem)
		}
		if res == maxNum {
			memo[amount] = -1
		} else {
			memo[amount] = res
		}
		return memo[amount]
	}

	return dp(coins, amount)
}

// 动态规划
// 322. 零钱兑换
// 自底向上
func coinChangeV(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}

	dp[0] = 0
	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
