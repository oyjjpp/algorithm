### 题目

76. 最小覆盖子串

#### 题目描述

给你一个字符串s、一个字符串t；返回s中涵盖t所有字符的最小子串。

#### 说明

如果s中不存在涵盖t所有字符的子串，则返回空字符串""。  
如果s中存在这样的子串，我们保证它是唯一的答案。  
涵盖t，则是不需要是有序的字串，只要保证是子串的排列接可以。  

#### 题目概述
就是说要在s中找到包含t中全部字母的一个子串，且这个子串一定是所有可能子串中最短的。  

### 案例

#### 示例一

```示例一
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
```

#### 示例二

```示例二
输入：s = "a", t = "a"
输出："a"
```

### 思路

#### 框架
1、我们在字符串S中使用双指针中的左右指针技巧，初始化left = right = 0，把索引左闭右开区间 [left, right) 称为一个「窗口」。  
2、我们先不断地增加right指针扩大窗口[left, right)，直到窗口中的字符串符合要求（包含了T中的所有字符）。  
3、此时，我们停止增加right，转而不断增加left指针缩小窗口[left, right)，直到窗口中的字符串不再符合要求（不包含T中的所有字符了）；同时，每次增加left，我们都要更新一轮结果。  
4、重复第2和第3步，直到right到达字符串S的尽头。  

#### 所需要的变量
定义一个map保存当前子串所有元素的个数：need  
定义一个窗口用于滑动保存当前满足条件的元素：window  
定义左右子针用户滑动窗口：left，right  
定义有效值，用于校验当前窗口是否满足字串:valid  
为求最小字串，定义对应的开始索引和长度：start，length   

### 参考

来源：力扣（LeetCode）  
链接：<https://leetcode-cn.com/problems/minimum-window-substring/>  
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。  

[labuladong的算法小抄](https://labuladong.gitbook.io/algo/di-ling-zhang-bi-du-xi-lie/hua-dong-chuang-kou-ji-qiao-jin-jie)
