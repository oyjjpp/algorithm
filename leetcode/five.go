// 给定一个字符串s，找到s中最长的回文子串,你可以假设s的最大长度为1000。
package leetcode

// 第一种思路：找出每一个子字符串，判断其是不是回文字符串
// 时间复杂度O(n^3)
func longestPalindrome(s string) string {
	length := len(s)
	if length == 0 {
		return ""
	}
	max := 0  //最大长度
	maxl := 0 // 左索引
	maxr := 0 // 右索引
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			// 校验所有子子字符串
			ret := isPalindromic(s, i, j)
			if ret {
				if max < (j - i + 1) {
					max = j - i + 1
					maxl = i
					maxr = j
				}
			}
		}
	}
	maxString := ""
	for i := maxl; i <= maxr; i++ {
		maxString += string(s[i])
	}
	return maxString

}

// isPalindromic
// 校验是否为回文子串
func isPalindromic(s string, l, r int) bool {
	for l < r {
		if s[l] == s[r] {
			l++
			r--
		} else {
			return false
		}

	}
	return true
}

// 第二种思路：遍历字符串，找出以"每个字符"为中心的回文字符串有多长，选最长的返回
// 时间复杂度O(n^2)
func longestPalindromeV2(s string) string {
	// 字符串中长度
	length := len(s)
	// babad
	getLen := func(i, j int) int {
		// 以s[i]s[j]为中心的最长回文字符串
		for i >= 0 && j < length {
			if s[i] == s[j] {
				i--
				j++
			} else {
				return j - i - 1
			}
		}
		return j - i - 1
	}
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	// 最大长度
	maxLen := 0
	// 开始位置
	maxStart := 0
	for i := 0; i < length; i++ {
		if tempLen := max(getLen(i, i+1), getLen(i, i)); tempLen > maxLen {
			maxLen = tempLen
			maxStart = i - (maxLen-1)/2
		}
	}

	maxString := ""
	for i := maxStart; i < maxStart+maxLen; i++ {
		maxString += string(s[i])
	}
	return maxString
}
