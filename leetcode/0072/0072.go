package leetcode

import (
    "log"
    "fmt"
)

// minDistance
// 72. 编辑距离
// 使用递归方式解决动态规划问题
func minDistance(word1 string, word2 string) int {
    // 备忘录
    memorandum := map[string]int{}
    
	var dp func(i, j int) int
	dp = func(i, j int) (rs int) {
		// TODO 验证
		defer func() {
			// log.Println(i, j, rs)
		}()
        
        // 备忘录解决重叠子问题
        key := fmt.Sprintf("%d%d", i, j)
        if data, ok := memorandum[key]; ok {
            log.Println(key)
            return data
        }
        
		// base case
		// 当一个字符串已经全部扫描完的情况，则直接累加另一个字符串剩余的长度
		if i == -1 {
			rs = j + 1
			return
		}
		if j == -1 {
			rs = i + 1
			return
		}

		// 状态转移
		// 两个字符串相等情况，则直接跳过
		if word1[i] == word2[j] {
			rs = dp(i-1, j-1)
            memorandum[key] = rs
			return
		}
		// 插入/删除/替换
		rs = min((dp(i, j-1) + 1), (dp(i-1, j) + 1), (dp(i-1, j-1) + 1))
        memorandum[key] = rs
		return
	}

	return dp(len(word1)-1, len(word2)-1)
}

// min
// 求最小值
func min(a, b, c int) int {
	minNums := a
	if b < minNums {
		minNums = b
	}
	if c < minNums {
		minNums = c
	}
	return minNums
}

// minDistanceV
// 编辑距离
// 使用DP table
func minDistanceV(word1 string, word2 string) int {
    m := len(word1)
    n := len(word2)
    // 创建DP数组存储切片
    dp := createSlice(m+1, n+1)
    
    // base case
    for i := 1; i <= m; i++ {
        dp[i][0] = i
    }
    
    for i := 1; i <= n; i++ {
        dp[0][i] = i
    }
    
    // 自底向上求解
    for i := 1 ;i <=m; i++ {
        for j := 1; j <= n; j ++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            }else{
                dp[i][j] = min(dp[i-1][j]+1,dp[i][j-1]+1,dp[i-1][j-1]+1)
            }
        }
    }
    log.Println(dp)
    return dp[m][n]
}

// createSlice
// 创建二维切片
func createSlice(m, n int) [][]int{
    data := make([][]int, m)
    for key:= range data{
        data[key] = make([]int, n)
    }
    return data
    
}
