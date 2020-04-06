package common

import (
	"unicode"
)

// 125、验证回文串
// 131、分割回文串
// 139、单词拆分
// 140、单词拆分||
// 208、实现Trie(前缀数)
// 212、单词搜索||
// 242、有效的字母异位词
// 387、字符串中的第一个唯一字符
// 344、反转字符串

// 回溯
// https://juejin.im/post/5de0f2ebf265da05d6510e5e#heading-0

// ReverseString
// 字符串反转
// @link https://juejin.im/post/5de4ebe6e51d4504d666fa77
func ReverseString(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// FirstUniqChar
// 字符串中的第一个唯一字符
// 思路 通过两次循环【两次循环索引是否一致】
// 第一次循环保存每个元素的最后出现索引
// 第二次循环验证，最后出现的索引与当前索引是否一致
func FirstUniqChar(s string) int {
	// 声明一个数组，记录每个元素的索引
	item := [26]int{}

	// 遍历当前字符串，并记录元素最后出现的索引
	for index, ch := range s {
		item[ch-'a'] = index
	}

	// 验证，第一次出现如果与最后一次出现一致，则返回当前索引
	for index, ch := range s {
		if index == item[ch-'a'] {
			return index
		} else {
			item[ch-'a'] = -1
		}
	}
	return -1
}

// IsAnagram
// 有效的字母异位词【两个字符串有同样的字母】
// @link https://juejin.im/post/5cff1601f265da1b94213fa4
func IsAnagram(s string, t string) bool {
	// 如果不相等 则不可能是字母异或位词
	if len(s) != len(t) {
		return false
	}
	// 初始化一个int数组记录26个字母出现的次数
	data := make([]int, 26)

	// S存在则增加一 T存在则索引减一
	for i := 0; i < len(s); i++ {
		data[s[i]-'a']++
		data[t[i]-'a']--
	}

	// 校验 所有26个字符位置都为0
	for i := 0; i < 26; i++ {
		if data[i] != 0 {
			return false
		}
	}
	return true
}

// FindWords
// 单词搜索 II
func FindWords(board [][]byte, words []string) []string {
	return []string{}
}

// isPalindrome
// 验证是否是有效回文字符串
func isPalindrome(s string) bool {
	// 判断是否合法的字符
	isValid := func(v rune) bool {
		return unicode.IsDigit(v) || unicode.IsLetter(v)
	}
	left, right := 0, len(s)-1

	for left < right {
		leftData, rightData := rune(s[left]), rune(s[right])

		// 当都不是有效字符时
		if !isValid(leftData) && !isValid(rightData) {
			left++
			right--
		} else if !isValid(leftData) {
			left++
		} else if !isValid(rightData) {
			right--
		} else if unicode.ToUpper(leftData) != unicode.ToUpper(rightData) {
			return false
		} else {
			left++
			right--
		}
	}
	return true
}

// partition
// 分割回文串
// link https://juejin.im/post/5cc1728a5188252d9053cda4
func partition(s string) [][]string {
	// 验证字符串是否为回文字符串
	isP := func(s string) bool {
		left, right := 0, len(s)-1
		for left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
		return true
	}

	// 存储最终结果
	rs := [][]string{}
	// 定义一个回溯函数,分割字符串进行校验
	// 参数
	// @param 需要切割的字符串
	// @param 当前索引切割状态
	DFS := func(string, []string) {}
	DFS = func(str string, data []string) {
		// 字符串长度为零
		if len(str) == 0 {
			rs = append(rs, append([]string{}, data...))
			return
		}

		for i := 1; i <= len(str); i++ {
			if isP(str[:i]) {
				DFS(str[i:], append(data, str[:i]))
			}
		}
	}

	DFS(s, []string{})
	return rs
}

// wordBreak
// 单词拆分
// 给定一个非空字符串s和一个包含非空单词列表的字典 wordDict，判定s是否可以被空格拆分为一个或多个在字典中出现的单词
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s)+1; j++ {
			for _, word := range wordDict {
				if s[i:j] == word && dp[i] == true {
					dp[j] = true
				}
			}
		}
	}
	return dp[len(s)]
}

// dp[i]：表示前i个元素是否能拆分成词典中的单词，故最终返回dp[len(s)]。
// 状态转移：初始化dp[0]=true，对每一个起始元素i遍历截止元素j（j范围[i+1, len(s)]），若前i个元素已经可以成功拆分且i到j也是一个可选单词，
// 则前j个元素也可以成功拆分，即d[j]置为true。

// StrtoInt
// 将一个字符串转换成一个整数，字符串不是一个合法的数值则返回0，要求不能使用字符串转换整数的库函数。
// 主要考察点
// 1、指针是否为空指针以及字符串是够为空字符串
// 2、字符串对于正负号的处理
// 3、输入值是否为合法值，即小于等于'9',大于等于'0'
// 4、int为32位，需要判断是否溢出
// 5、使用错误标志，区分合法值0和非合法值
func StrtoInt(str string) int {
	if str == "" || len(str) == 0 {
		return 0
	}

	// 判断是否为负数
	isNegative := str[0] == '-'

	// 保存结果
	ret := 0
	for i := 0; i < len(str); i++ {
		cur := str[i]
		// 当第一个字符出现“+”，“-”则继续
		if i == 0 && (cur == '+' || cur == '-') {
			continue
		}
		// 验证是否为非法输入
		if cur < '0' || cur > '9' {
			return 0
		}
		// 移位
		ret = ret*10 + int(cur-'0')
	}

	if isNegative {
		return -ret
	} else {
		return ret
	}
}
