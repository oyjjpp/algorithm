### 题目

516. 最长回文子序列

### 题目描述

给定一个字符串 s ，找到其中最长的回文子序列，并返回该序列的长度。可以假设 s 的最大长度为 1000 。

### 案例

#### 示例1

```golang
输入：“bbbab”
输出：4
解释：一个可能的最长回文子序列为 "bbbb"。
```

#### 示例2

```golang
输入：“cbbd”
输出：2
解释：一个可能的最长回文子序列为 "bb"。
```

### 提示

>1 <= s.length <= 1000  
>s 只包含小写英文字母  

### 思路

一旦涉及子序列和最值问题，几乎可以肯定考察的是动态规划技巧，时间复杂度为O(n^2)  

动态规划问题，主要条件base case ，状态转移，最优结构 DP数组（函数）  
base case:dp[i][j] = 1[i==j], dp[i][j]=0[i<j]
状态转移:

```golang
 if s[i]==s[j] {
  dp[i][j] = d[i+1][j-1]+2
 } else {
  dp[i][j] = max(d[i+1][j], d[i][j-1])
 }
```

最优结构：dp[0][n-1]

### 代码

### 参考
