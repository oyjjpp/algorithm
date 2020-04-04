package common

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
// 思路 通过两次循环
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
