package leetcode

import (
	"testing"
)

func TestSlice(t *testing.T) {
	data := []int{1, 2, 3}
	t.Log(data[len(data)-1])
}

func TestFindNumberOfLIS(t *testing.T) {
	data := []int{1, 3, 5, 4, 7}
	rs := findNumberOfLIS(data)
	t.Log(rs)
}

func TestFindMaxLengthOfLIS(t *testing.T){
    data := []int{1, 4, 3, 4, 2}
    rs := findMaxLengthOfLIS(data)
    t.Log(rs)
    rs2 := findMaxLengthOfLISV2(data)
    t.Log(rs2)
}
