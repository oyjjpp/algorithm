package common

// 爬楼梯问题
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 使用最小花费爬楼梯
func minCostClimbingStairs(cost []int) int {
	f0 := 0
	f1 := 0
	f2 := 0
	for i := len(cost) - 1; i >= 0; i-- {
		if f1 < f2 {
			f0 = cost[i] + f1
		} else {
			f0 = cost[i] + f2
		}
		f2 = f1
		f1 = f0
	}
	if f1 < f2 {
		return f1
	} else {
		return f2
	}
}
