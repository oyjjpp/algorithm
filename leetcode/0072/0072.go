package leetcode

// minDistance
// 72. 编辑距离
// 使用递归方式解决动态规划问题
func minDistance(word1 string, word2 string) int {
    var dp func(i, j int)int
    dp = func(i, j int) int {
        // base case
        // 当一个字符串已经全部扫描完的情况，则直接累加
        if i == -1 {
            return i+1
        }
        if j == -1 {
            return j+1
        }

        // 状态转移
        // 两个字符串相等情况，则直接跳过
        if word1[i] == word2[j]{
            return dp(i-1, j-1)
        } 
        // 插入/删除/替换
        return min((dp(i, j-1) + 1), (dp(i-1, j) + 1), (dp(i-1, j-1) + 1))
    }
    
    return dp(len(word1)-1, len(word2)-1)
}

// min
// 求最小值
func min(a, b, c int) int {
    minNums := a
    if b < minNums{
        minNums = b
    }
    if c < minNums{
        minNums = c
    }
    return minNums
}
