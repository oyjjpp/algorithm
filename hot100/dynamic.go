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

// 状态 、 选择
// 状态 -> 选择 （穷举）

func maxProfit(prices []int) int {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	k := 2
	days := len(prices)
	data := make([][][]int, days)
	for i := range data {
		data[i] = make([][]int, k+1)
		for j := range data[i] {
			data[i][j] = make([]int, 2)
		}
	}

	for i := 0; i < days; i++ {
		if i == 0 {
			data[i][k][0] = 0
			data[i][k][1] = -prices[i]
			continue
		}

		data[i][k][0] = max(data[i-1][k][0], data[i-1][k][1]+prices[i])
		data[i][k][1] = max(data[i-1][k][1], data[i-1][k-1][0]-prices[i])
	}
	return data[days-1][k][0]
}

// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码还未经过力扣测试，仅供参考，如有疑惑，可以参照我写的 java 代码对比查看。

// 原始版本
func maxProfit_k_2(prices []int) int {

	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	max_k := 2 // 最大可交易次数
	n := len(prices)
	dp := make([][][]int, n) // i为天数，k为当前第几次交易，0表示不持有股票，1表示持有股票
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, max_k+1)
		for k := 0; k < max_k+1; k++ {
			dp[i][k] = make([]int, 2)
		}
	}
	for i := 0; i < n; i++ {
		for k := max_k; k >= 1; k-- {
			if i-1 == -1 {
				// 处理 base case
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	// 穷举了 n × max_k × 2 个状态，正确。
	return dp[n-1][max_k][0] //返回最大利润
}

func maxProfit_v(prices []int) int {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	k := 2
	n := len(prices)
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, k+1) // 构建 3 维 DP 数组
		for j := range dp[i] {
			dp[i][j] = make([]int, 2) // 初始化 DP 数组
		}
	}

	for i := 0; i < n; i++ {
		for j := k; j >= 1; j-- {
			if i-1 == -1 {
				// 处理 base case
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
				continue
			}
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}

	return dp[n-1][k][0]
}
