package leetcode

import "log"

// minWindow
// 最小覆盖子串
func minWindow(s string, t string) string {
	// 初始化计数器，分别记录【窗口】中字符的出现次数和T中字符串出现次数
	window := map[byte]int{}
	need := map[byte]int{}
	for i := 0; i < len(t); i++ {
		cur := t[i]
		need[cur]++
	}

	maxNum := 1<<31 - 1

	// 初始化窗口的两端，区间[left, right)
	left, right := 0, 0

	// 表示窗口中满足need条件的字符个数
	// 如果valid和len(need)的大小相同，则说明窗口已满足条件，已经完全覆盖了串T
	valid := 0

	// @TODO 记录最小覆盖子串的起始索引及长度
	start, length := 0, maxNum
	log.Println(right, left, start, length)
	for right < len(s) {
		cur := s[right]

		// 进行窗口数据更新
		if _, ok := need[cur]; ok {
			window[cur]++

			// 表示窗口中满足need条件的字符个数
			if window[cur] == need[cur] {
				valid++
			}
		}

		// 判断窗口是否需要收缩
		for valid == len(need) {
			// 在这里更新最小覆盖子串
			if (right - left) < length {
				start = left
				length = right - left
			}

			// 将移出窗口的字符
			delS := s[left]
			left++
			if _, ok := need[delS]; ok {
				if window[delS] == need[delS] {
					valid--
				}
				window[delS]--
			}
		}
	}
	// 返回最小覆盖子串
	if length == maxNum {
		return ""
	} else {
		return s[start:length]
	}
}
