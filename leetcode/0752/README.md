### 题目

(752). 打开转盘锁

#### 题目描述

你有一个带有四个圆形拨轮的转盘锁；每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' ；
每个拨轮可以自由旋转：例如把'9'变为'0'，'0' 变为'9' ；每次旋转都只能旋转一个拨轮的一位数字。  
锁的初始数字为'0000' ，一个代表四个拨轮的数字的字符串。  
列表deadends包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。  
字符串target代表可以解锁的数字，你需要给出最小的旋转次数，如果无论如何不能解锁，返回-1。  

### 示例

#### 示例一
```conf
输入：["0201","0101","0102","1212","2002"] "0202"
输出：6
```

### 思路一
使用BFS  
明确起点和终点    

起点：”0000“
终点：target

当前节点扩散的点：每个位置向上拨动或者向下拨动共有八种情景

死亡数字：遇到死亡数值直接掠过

优化：为防止重复遍历已经遍历的点，可以将遍历的节点存储起来，加入扩散的点时进行校验

#### 代码

```golang

```


### 参考

来源：力扣（LeetCode）  
链接：<https://leetcode-cn.com/problems/open-the-lock/>  
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。  

[labuladong的算法小抄](https://labuladong.gitbook.io/algo/di-ling-zhang-bi-du-xi-lie/bfs-kuang-jia)