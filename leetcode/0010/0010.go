package leetcode

import (
    "fmt"
)

// isMatch
// 正则匹配
func isMatch(s string, p string) bool {
    // 记录状态（i,j） 消除重叠子问题
    memo := map[string]bool{}
    
    var dp func(s, p *string, i, j int) bool 
    
    dp = func(s, p *string, i, j int) bool {
        m, n := len(*s), len(*p)
        
        // base case
        if j==n {
            return j==n
        }
        
        if i==m {
            // 正则表达是剩余的字符
            // TODO 如果只剩下. 匹配成“”
            if (n-j) % 2 ==1 {
                return false
            }
            // x*y*z*
            for ;j+1 <n;j+=2 {
                if *p[j+1] !='*'{
                    return false
                }
            }
            return true
        }
        key := fmt.Sprintf("%d-%d",i,j)
        if rs,ok := memo[key]; ok {
            return rs
        }
        res := false
        // 两个字符相同或者正则表达式是"."
        if *s[i] == *p[j] || *p[j] == '.'{
                if j < n-1 && *p[j+1] == '*' {
                    res = dp(s, p, i, j+2) || dp(s, p, i+1, j)
                } else {
                    res = dp(s, p, i+1, i+1)
                } 
        } else {
            if j <n-1 && *p[j+1]=='*' {
                res = dp(s, p, i, j+2)
            }else{
                return false
            }
        }
        memo[key] = res
        return res
    }
    return dp(&s, &p, 0, 0)
}
