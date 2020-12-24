package leetcode

import (
	"log"
)

// minInsertions
// 让字符串成为回文串的最少插入次数
func minInsertions(s string) int {
	n := len(s)
	if n <= 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	// base case
	dp := makeSlice(n, n)

	// 状态转移
	// 从下向上
	for i := n - 2; i >= 0; i-- {
		// 从左到右遍历
		for j := i + 1; j < n; j++ {
			log.Println(n, i, j)
			// 根据s[i]和s[j]进行状态转移
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[0][n-1]
}

func makeSlice(m, n int) [][]int {
	data := make([][]int, m)
	for key := range data {
		data[key] = make([]int, n)
	}
	return data
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
