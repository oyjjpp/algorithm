package leetcode

import (
    "log"
)

// longestPalindromeSubseq
// 最长回文子序列
func longestPalindromeSubseq(s string) int {
    length := len(s)
    if length == 0 {
        return 0
    }
    if length == 1 {
        return 1
    }
    
    // 定义 base case
    dp := makeSlice(length, length)
    for i:=0; i< length; i++{
        dp[i][i] = 1
    }
    log.Println(dp)
    
    // 状态转移方程
    for i := length-2; i>=0; i-- {
        for j := i+1; j<length;j++ {
            log.Println(s[i] == s[j], i, j)
            if s[i] == s[j] {
                dp[i][j] = dp[i+1][j-1]+2
            } else {
                dp[i][j] = max(dp[i+1][j], dp[i][j-1])
            } 
        }
    }
    return dp[0][length-1]
}

// makeSlice
// 创建二维切片
func makeSlice(m, n int) [][]int{
    data := make([][]int, m)
    for key := range data{
        data[key] = make([]int, n)
    }
    return data
}

func max(a, b int) int{
    if a > b {
        return a
    }
    return b
}
