package leetcode

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	rs := lengthOfLongestSubstring("pwwkew")
	t.Log(rs)

	rs = lengthOfLongestSubstring("bbbbb")
	t.Log(rs)

	rs = lengthOfLongestSubstring("abcabcbb")
	t.Log(rs)

	param := map[string]int{
		"abcabcbb": 3,
		"bbbbb":    1,
		"pwwkew":   3,
		"dedf":     3,
		" ":        1,
	}
	for k, v := range param {
		rs := lengthOfLongestSubstring(k)

		if rs != v {
			t.Errorf("expect %d, result %d\n", v, rs)
		}
	}
}
