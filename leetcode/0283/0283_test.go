package leetcode

import "testing"

func TestMoveZeroes(t *testing.T) {
	data := []int{0, 1, 0, 3, 12}
	moveZeroes(data)
	t.Log(data)
}
