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
