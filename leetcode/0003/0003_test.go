package leetcode

import (
    "testing"
)

func TestLengthOfLongestSubstring(t *testing.T){
    rs := lengthOfLongestSubstring("pwwkew")
    t.Log(rs)

    rs = lengthOfLongestSubstring("bbbbb")
    t.Log(rs)

    rs = lengthOfLongestSubstring("abcabcbb")
    t.Log(rs)
}
