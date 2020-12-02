package leetcode

// findAnagrams
// 438. 找到字符串中所有字母异位词
func findAnagrams(s, p string) []int {
    // 初始化计数器，分别记录【窗口】中字符的出现次数和T中字符串出现次数
	window := map[byte]int{}
	need := map[byte]int{}
	for i := 0; i < len(p); i++ {
		cur := p[i]
		need[cur]++
	}
	// 初始化窗口的两端，区间[left, right)
	left, right := 0, 0

	// 表示窗口中满足need条件的字符个数
	// 如果valid和len(need)的大小相同，则说明窗口已满足条件，已经完全覆盖了字符串t
	valid := 0
    
    // 保存结果
    data := []int{}
    
    // 结束条件：存在满足条件的节点即可
	for right < len(s) {
        // 即将移入窗口的值
		cur := s[right]
        right++
        
		// 进行窗口数据更新
        // 首先校验是否在T串
		if _, ok := need[cur]; ok {
			window[cur]++

			// 表示窗口中满足need条件的字符个数
			if window[cur] == need[cur] {
				valid++
			}
		}

		// 判断窗口是否需要收缩：窗口大小大于len(need)时，应为排列，显然长度应该是一样的。
		for (right - left) >= len(p) {
            // 结束条件：在这里判断是否找到了合法的子串,如果合法则将左指针存入结果中
            if valid == len(need){
                data = append(data, left)
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
    return data
}
