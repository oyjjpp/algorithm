package leetcode

import (
	"log"
)

// longestCommonSubsequence
// 1143. 最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	w := len(text1)
	h := len(text2)

	// base case
	// 初始化
	dp := MakeIntSlice(w+1, h+1)
	log.Println(dp)

	for i := 1; i <= w; i++ {
		for j := 1; j <= h; j++ {
			// 状态转移
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[w][h]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MakeSlice
// 创建二维切片
func MakeIntSlice(row, column int) [][]int {
	data := make([][]int, row)
	for index := range data {
		data[index] = make([]int, column)
	}
	return data
}
