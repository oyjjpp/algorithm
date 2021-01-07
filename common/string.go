package common

import (
	"math"
	"strconv"
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

// isPalindromeStr
// 验证是否是有效回文字符串
func isPalindromeStr(s string) bool {
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

// "123124646664124” + “6775685684”
func stringAdd(str1, str2 string) string {
	sum := 0
	num1 := len(str1)
	num2 := len(str2)
	res := []int{}
	for num1 > 0 || num2 > 0 || sum > 0 {
		total := sum
		if num1 > 0 {
			num1--
			cur := str1[num1]
			total = total + int(cur-'0')
		}

		if num2 > 0 {
			num2--
			total = total + int(str2[num2]-'0')
		}

		// 考虑进位问题
		sum = total / 10
		res = append(res, int(total%10))
	}
	resut := ""
	for _, value := range res {
		cur := strconv.Itoa(value)
		resut = cur + resut
	}
	return resut
}

// 字符串转换成整数
func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	// 去掉字符串前面空格
	stripLeadingWhitespace := func(str string) string {
		result := ""
		isLeading := true
		for _, char := range str {
			if isLeading {
				if char != ' ' {
					isLeading = false
					result = result + string(char)
				}
			} else {
				result = result + string(char)
			}
		}
		return result
	}
	strippedStr := stripLeadingWhitespace(str)
	sign := 1
	result := 0
	for i, char := range strippedStr {
		if i == 0 {
			if char == '-' {
				sign = -1
			} else if char == '+' {
				sign = 1
			} else if int(char-'0') >= 0 && int(char-'0') <= 9 {
				result = result*10 + int(char-'0')
			} else {
				return result
			}
		} else {
			if int(char-'0') < 0 || int(char-'9') > 9 {
				break
			}
			// 向下溢出
			if sign == -1 && result*10+sign*int(char-'0') < math.MinInt32 {
				return math.MinInt32
			}
			// 向上溢出
			if sign == 1 && result*10+int(char-'0') > math.MaxInt32 {
				return math.MaxInt32
			}
			result = result*10 + sign*int(char-'0')
		}
	}
	return result
}

// stripLeadingWhitespace
// 去掉开头的空白字符串
func stripLeadingWhitespace(str string) string {
	result := ""
	isLeading := true
	for _, char := range str {
		if isLeading {
			if char != ' ' {
				isLeading = false
				result = result + string(char)
			}
		} else {
			result = result + string(char)
		}
	}
	return result
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	// 最长的公共前缀的长度肯定小于等于数组中最短的元素
	// 所以从这个元素开始当基准
	short := findShortestInArray(strs)
	// 早发现早治疗
	if len(short) == 0 {
		return ""
	}
	// 遍历这个最短的每个位置元素，用来判断是不是相等
	for i, v := range short {
		// 要判断多少次，取决于数组strs中有多少个元素，所以用的len(strs)
		for j := 0; j < len(strs); j++ {
			// 数组的第j个元素的第i个位置不等于我们的short的第i个位置的元素
			// 写成strs[j][i] 是为了和short里面的每个元素一一对应比较
			if strs[j][i] != byte(v) {
				// 到了第[j][i]个没有匹配上，那么就说明之前的都匹配上了，所以直接返回[j][:i]
				return strs[j][:i]
			}
		}
	}
	// 遍历完short了，说明short就是最长的，直接返回
	return short
}

func findShortestInArray(s []string) string {
	// 空字符数组返回空
	if len(s) == 0 {
		return ""
	}
	// 临时定义最短为数组第一个
	shortest := s[0]
	// 遍历数组每个元素
	for _, v := range s {
		// 找到当前小于res
		if len(v) < len(shortest) {
			// 看看是否是空的，空的说明数组中有空字符，所以最长公共前缀肯定为空
			if len(v) == 0 {
				return ""
			}
			// 替换当前最小为当前遍历到的元素
			shortest = v
		}
	}
	return shortest
}
