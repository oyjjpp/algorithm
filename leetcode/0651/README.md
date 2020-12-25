### 题目

651 四键键盘

### 题目描述

假设有一个特殊的键盘，上面只有四个键，他们分别是：  
A键：在屏幕上显示一个A  
Ctrl-A键：选中整个屏幕  
Ctrl-C键：将选中的区域复制到缓冲区  
Ctrl-V键：将缓冲区的内容输出到光标所在的屏幕位置

现在要求只能进行N次操作，请计算屏幕上最多能显示多少个A？

### 案例

#### 示例1

```golang
输入：3
输出：3
因为连续按3次A键是最优的方案
```

#### 示例2

```golang
输入：7
输出：9
最优的操作序列：A，A，A，Ctrl-A，Ctrl-C，Ctrl-V，Ctrl-V
```

### 思路

状态
>当前可以操作的次数N  
>屏幕上A的数量sumA  
>缓冲区有多少cacheA

base case：N=0

状态转移：  
按下A：dp(N-1, sumA+1, cacheA)  
按下Ctrl+V：dp(N-1, sumA+cacheA, cacheA)
按下Ctrl+A，Ctrl+C：dp(N-2, sumA, sumA)  
