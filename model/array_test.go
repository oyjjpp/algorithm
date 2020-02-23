package model

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	rs := TwoSum(nums, target)
	if len(rs) != 2 {
		t.Errorf("希望结果长度为:2，实际结果长度:%d", len(rs))
	}
	if rs[0] != 0 || rs[1] != 1 {
		t.Errorf("希望索引位置[0,1]，实际结果索引位置[%d,%d]", rs[0], rs[1])
	}
}

func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TestTwoSum(&testing.T{})
	}
}

func TestTwoSumHash(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	rs := TwoSumHash(nums, target)
	if len(rs) != 2 {
		t.Errorf("希望结果长度为:2，实际结果长度:%d", len(rs))
	}
	if rs[0] != 0 || rs[1] != 1 {
		t.Errorf("希望索引位置[0,1]，实际结果索引位置[%d,%d]", rs[0], rs[1])
	}
}

func BenchmarkTwoSumHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TestTwoSumHash(&testing.T{})
	}
}
