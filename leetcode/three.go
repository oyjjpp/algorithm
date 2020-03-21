/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
*/
package leetcode

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	// 使用map存储 空间ascii长度
	data := make(map[rune]int, 128)
	// 索引位置
	index := 0
	for key, value := range s {
		// 校验是否存在过
		if pos, ok := data[value]; ok {
			// 替换索引
			if pos > index {
				index = pos
			}
		}
		// 防止字符串长度为1时,key-index=0
		if key+1-index > maxLen {
			maxLen = key + 1 - index
		}
		data[value] = key + 1
	}
	return maxLen
}
