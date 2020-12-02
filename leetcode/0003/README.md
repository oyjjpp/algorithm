### 题目
3. 无重复字符的最长子串

#### 题目描述
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

### 示例

#### 示例一
```config
输入: s = "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
```

#### 示例二
```config
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
```

#### 示例三
```config
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
```

### 思路
1、使用一个窗口记录浏览过的元素  
2、初始化两个指针left,right;使用right先扫描元素  
3、碰到重复的元素调整left，去掉窗口中之前记录的元素  
4、初始时初始化一个子串长度，随着left和right更新时，逐步更新  

### 代码
```golang
```

### 同类——双指针
[算法—leetcode—438—找到字符串中所有字母异位词](https://juejin.cn/post/6901683129849741320)  
[算法—leetcode—567—字符串的排列](https://juejin.cn/post/6901209346459074573)  
[算法—leetcode—76—最小覆盖子串](https://juejin.cn/post/6901184530648924167)  

### 参考
来源：力扣（LeetCode）  
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters  
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。  

[labuladong的算法小抄](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)
