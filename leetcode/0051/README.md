### 题目

（51.）N 皇后

#### 题目描述

n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

#### 提示

皇后彼此不能相互攻击，也就是说：任何两个皇后都不能处于同一条横行、纵行或斜线上。  

### 示例

#### 示例一

```conf
输入：4
输出：[
 [".Q..",  // 解法 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // 解法 2
  "Q...",
  "...Q",
  ".Q.."]
]
解释: 4 皇后问题存在两个不同的解法。
```

### 思路

#### 框架
1、路径：也就是已经做出的选择；  
2、选择列表：当前可以做的选择；  
3、结束条件：到达决策树底层，无法再做选择的条件。  

#### N皇后问题对应框架的解释
路径：board中小于row的那些行都已经成功放置了皇后  
选择列表：第row行的所有列都是放置皇后的选择  
结束条件：row超过board的最后一行  

#### 准备条件
1、初始化一个棋盘的二为数组
2、定义一个校验是否有效位置的函数isValid
3、按框架思路书写循环代码

### 框架

```框架
for 选择 in 选择列表:
    # 做选择
    将该选择从选择列表移除
    路径.add(选择)
    backtrack(路径, 选择列表)
    # 撤销选择
    路径.remove(选择)
    将该选择再加入选择列表
```

### 代码

```golang

```

### 参考

来源：力扣（LeetCode）  
链接：<https://leetcode-cn.com/problems/n-queens//>  
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。  

[labuladong的算法小抄](https://labuladong.gitbook.io/algo/)