### 题目

1312. 让字符串成为回文串的最少插入次数

### 描述

给你一个字符串 s ，每一次操作你都可以在字符串的任意位置插入任意字符。  
请你返回让 s 成为回文串的 最少操作次数 。  
「回文串」是正读和反读都相同的字符串。  

### 示例

#### 案例一

```golang
输入：s = "zzazz"
输出：0
解释：字符串 "zzazz" 已经是回文串了，所以不需要做任何插入操作。
```

#### 案例二

```golang
输入：s = "mbadm"
输出：2
解释：字符串可变为 "mbdadbm" 或者 "mdbabdm" 。
```

#### 案例三

```golang
输入：s = "leetcode"
输出：5
解释：插入 5 个字符后字符串变为 "leetcodocteel" 。
```

#### 案例四

```golang
输入：s = "g"
输出：0
```

#### 案例五

```golang
输入：s = "no"
输出：1
```

### 思路

最少插入次数，求最值还是使用动态规划思路
需要明确各个条件
base case : dp[i][j]=0 {i==j}

状态转移方程

```golang
 if s[i] == s[j] {
  dp[i][j] = dp[i+i][j-1]
 } else {
  dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1
 }
```

最优解:dp[0][n-1]

### 代码

### 参考

来源：力扣（LeetCode）  
链接：<https://leetcode-cn.com/problems/minimum-insertion-steps-to-make-a-string-palindrome>  
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。  
