package leetcode

import "testing"

func TestSuperEggDrop(t *testing.T) {
	// 1 2 2
	// 2 6 3
	// 3 14 4
	// 4 2000
	rs := superEggDropV2(4, 2000)
	t.Log(rs)
}
