### 题目 
53. 最大子序和  连续子数组的最大和  

### 题目描述

输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。  
要求时间复杂度为O(n)。

### 案例
```golang
输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
```

### 思路
1,使用动态规划思想  
2,寻找基本状态/状态转移/Db table  
3,本题是连续的整数组成的子数组  

### 优化
因为当前状态转移dp[i]至于dp[i-1]有关系，则我们可以通过状态压缩，同时求出最大值
maxNums := dp[0]  
maxNums = max(maxNums, dp[i]) 


