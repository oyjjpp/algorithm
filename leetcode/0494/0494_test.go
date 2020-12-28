package leetcode

import "testing"

func TestFindTargetSumWays(t *testing.T) {
	data := []int{1, 1, 1, 1, 1}
	rs := findTargetSumWaysV2(data, 3)
	t.Log(rs)
}
