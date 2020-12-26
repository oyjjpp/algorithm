package leetcode

// climbStairs
// 70.爬楼梯
func climbStairs(n int) int {
	// n 为正整数
	if n < 2 {
		return 1
	}
	dp := make([]int, n+1)
	// base case
	dp[1] = 1
	dp[2] = 2
	// 状态转移
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
