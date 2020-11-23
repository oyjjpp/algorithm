#

## algorithm

Algorithm program of record learning

从整体到细节，自顶向下，从抽象到具体的框架思维是通用的

### 数据结构的存储方式

数组（顺序存储）和链表（链式存储）

数组
>数组是连续存储，可以随机访问，通过索引快速定位元素，相对节约存储空间
>因为是连续存储，针对扩容操作，需要重新开辟一块连续存储空间，然后将元素都复制过去，中间插入，删除操作都需要将插入位子后续的元素进行移动

链表
>元素不连续，是通过一个指针指向下一个元素，不存在连续存储扩容的问题，而且针对插入和删除操作只需要更改指针即可
>因为存储不连续，不能根据索引定位某一元素，所以不能随机访问，而且每个元素除存储值之外还要存储下一个元素指针，会相对消耗更多时间

### 数据结构的基本操作

对于任何数据结构，其基本操作无非遍历 + 访问，再具体一点就是：增删查改。  

访问的方式主要是线性和非线性两种方式  
>线性访问为常见的for/while迭代的方式  
>非线性是递归的代表  

### 动态规划问题

动态规划问题一般就是求最值，求解动态规划的核心问题是穷举；因为要求最值，肯定要把所有可行的答案穷举出来，然后在其中找最值；  
动态规划的穷举存在“重叠子问题”，暴力穷举的话效率很低，所以需要“备忘录”或“DP table”来优化穷举过程，避免不必要的计算；
而且，动态规划问题一定会具备“最优子结构”，才能通过子问题的最值得到原问题的最值。

动态规划问题解题三要素  
重叠子问题、最优子结构、状态转移方程（难点）

明确base case -> 明确【状态】 -> 明确【选择】 -> 定义DP数组/函数的含义

递归算法的时间复杂度？就是用子问题个数乘以解决一个子问题需要的时间

动态规划问题最困难的就是写出这个暴力解，即状态转移方程

要符合「最优子结构」，子问题间必须互相独立

自顶向下：递归

自下向上：迭代(DP table)  
从最小状态列举到最终状态值

### 回溯问题

解决一个回溯问题，实际上就是一个决策树的遍历过程。

1、路径：也就是已经做出的选择。  
2、选择列表：也就是你当前可以做的选择。  
3、结束条件：也就是到达决策树底层，无法再做选择的条件。  

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

### 参考

[labuladong的算法小抄](https://labuladong.gitbook.io/algo/)
