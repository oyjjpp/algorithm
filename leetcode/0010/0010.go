package leetcode

import (
	"encoding/json"
	"fmt"
	"log"
)

// isMatch
// 正则匹配
func isMatch(s string, p string) bool {
	// 记录状态（i,j） 消除重叠子问题
	memo := map[string]bool{}

	var dp func(s, p *string, i, j int) bool
	dp = func(s, p *string, i, j int) bool {
		log.Println("当前索引", i, j)
		// 字符串的长度
		m, n := len(*s), len(*p)

		// base case
		// 正则表达式已经匹配完成
		if j == n {
			log.Println(j)
			return i == m
		}

		// 字符串已经匹配完成
		if i == m {
			log.Println("字符串已经匹配完成", n, j)
			// 正则表达是剩余的字符
			// TODO 如果只剩下. 匹配成“”
			if (n-j)%2 == 1 {
				return false
			}
			// x*y*z*
			for ; j+1 < n; j += 2 {
				if (*p)[j+1] != '*' {
					return false
				}
			}
			return true
		}
		// 备忘录的key
		key := fmt.Sprintf("%d-%d", i, j)
		if rs, ok := memo[key]; ok {
			return rs
		}
		if rs, err := json.Marshal(memo); err == nil {
			log.Println("备忘录", string(rs))
		}
		// "aab" "c*a*b" true
		res := false
		// 两个字符相同或者正则表达式是"."
		if (*s)[i] == (*p)[j] || (*p)[j] == '.' {
			if j < n-1 && (*p)[j+1] == '*' {
				res = dp(s, p, i, j+2) || dp(s, p, i+1, j)
			} else {
				res = dp(s, p, i+1, j+1)
			}
		} else {
			// 两个字符串不相同
			if j < n-1 && (*p)[j+1] == '*' {
				res = dp(s, p, i, j+2)
			} else {
				return false
			}
		}
		memo[key] = res
		return res
	}
	return dp(&s, &p, 0, 0)
}
