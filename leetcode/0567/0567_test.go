package leetcode

import (
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	rs := checkInclusion("ab", "eidbaooo")
	t.Log(rs)

	rs = checkInclusion("ab", "eidboaoo")
	t.Log(rs)

	rs = checkInclusion("abcdxabcde", "abcdeabcdx")
	t.Log(rs)
}
