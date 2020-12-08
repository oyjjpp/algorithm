package leetcode

import (
	"testing"
)

func TestSlice(t *testing.T) {
	data := []int{1, 2, 3}
	t.Log(data[len(data)-1])
}

func TestFindMaxLengthOfLIS(t *testing.T) {
	data := []int{1, 3, 5, 4, 7}
	rs := findMaxLengthOfLIS(data)
	t.Log(rs)
}
