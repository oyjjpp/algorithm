### 题目
438. 找到字符串中所有字母异位词

#### 题目描述
给定一个字符串 s 和一个非空字符串p，找到s中所有是p的字母异位词的子串，返回这些子串的起始索引。  
字符串只包含小写英文字母，并且字符串s和p的长度都不超过20100。

#### 说明
字母异位词指字母相同，但排列不同的字符串。  
不考虑答案输出的顺序。  

### 案例

#### 示例1
```config
输入:
s: "cbaebabacd" p: "abc"

输出:
[0, 6]

解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
```

#### 示例2
```config
输入:
s: "abab" p: "ab"

输出:
[0, 1, 2]

解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。
```

### 思路
#### 框架
```框架
/* 滑动窗口算法框架 */
void slidingWindow(string s, string t) {
    unordered_map<char, int> need, window;
    // 初始化字串所有元素的个数
    for (char c : t) need[c]++;

	// 初始化左右指针
    int left = 0, right = 0;
    // 有效性的个数
    int valid = 0;
    while (right < s.size()) {
        // c 是将移入窗口的字符
        char c = s[right];
        // 右移窗口
        right++;
        // 进行窗口内数据的一系列更新
        ...

        /*** debug 输出的位置 ***/
        printf("window: [%d, %d)\n", left, right);
        /********************/

        // 判断左侧窗口是否要收缩
        while (window needs shrink) {
            // d 是将移出窗口的字符
            char d = s[left];
            // 左移窗口
            left++;
            // 进行窗口内数据的一系列更新
            ...
        }
    }
}
```


### 代码

### 参考
来源：力扣（LeetCode）  
链接：https://leetcode-cn.com/problems/find-all-anagrams-in-a-string  
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。  

[labuladong的算法小抄](https://labuladong.gitbook.io/algo/di-ling-zhang-bi-du-xi-lie/hua-dong-chuang-kou-ji-qiao-jin-jie#san-zhao-suo-you-zi-mu-yi-wei-ci)
