package leetcode

import "testing"

func TestMaxCoins(t *testing.T) {
	data := []int{3, 1, 5, 8}
	rs := maxCoinsV2(data)
	t.Log(rs)
}
