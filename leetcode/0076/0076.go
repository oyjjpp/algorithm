package leetcode

import "log"

// minWindow
// 最小覆盖子串
// @param s 原字符串
// @param t 子字符串
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

	// 表示窗口中满足need的字符个数
	// 如果valid和len(need)的大小相同，则说明窗口已满足条件，已经完全覆盖了字符串t
	valid := 0

	// @TODO 记录最小覆盖子串的起始索引及长度，计算最终结果使用
	start, length := 0, maxNum
	log.Println(right, left, start, length)
    
    // 结束条件：right到达s的尽头
	for right < len(s) {
        // 即将移入窗口的值
		cur := s[right]
        right++
        
		// 进行窗口数据更新
        // 首先校验当前查找的字符是否在字串中
		if _, ok := need[cur]; ok {
			window[cur]++

			// 表示窗口中满足need条件的字符个数
			if window[cur] == need[cur] {
				valid++
			}
		}
  
		// 判断窗口是否需要收缩：valid==len(need)代表当前窗口已经满足need
		for valid == len(need) {
			// 在这里更新最小覆盖子串：最终结果产生位置
			if (right - left) < length {
				start = left
				length = right - left
			}

			// 将移出窗口的字符
			delS := s[left]
			left++
			if _, ok := need[delS]; ok {
                // 既要移除窗口的值，也要移除验证的值
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
        log.Println(start, length)
		return s[start:start+length]
	}
}
