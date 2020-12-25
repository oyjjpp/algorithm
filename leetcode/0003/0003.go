package leetcode

// lengthOfLongestSubstring
// 3. 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	window := map[byte]int{}
	// 左右指针
	left, right := 0, 0
	// 保存结果
	length := 0
	for right < len(s) {
		// 更新窗口
		cur := s[right]
		right++
		window[cur]++

		for window[cur] > 1 {
			delCur := s[left]
			left++
			window[delCur]--
		}
		length = max(length, right-left)
	}
	return length
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstringV2(s string) int {
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
