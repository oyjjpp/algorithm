### 题目

72. 编辑距离

### 题目描述

给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

>插入一个字符
>删除一个字符
>替换一个字符

### 示例

#### 示例一

```golang
输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
```

#### 示例二

```golang
输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')
```

### 思路

动态规划问题，使用递归方式解决  
1、明确base case  
当一个字符串已经全部扫描完的情况，则直接累加另一个字符串剩余的长度  
2、状态转移  
两个字符串相等情况，则直接跳过  
否则求插入/删除/替换中最小的方式  
3、使用备忘录解决重复子问题

### 参考

来源：力扣（LeetCode）
链接：<https://leetcode-cn.com/problems/edit-distance>  
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。  
