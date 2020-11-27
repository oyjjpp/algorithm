package leetcode

import (
    "testing"
)

func TestSearchRange(t *testing.T){
    data := []int{5,7,7,8,8,10}
    rs := searchRange(data, 8)
    t.Log(rs)
    
    rs1 := searchRange(data, 6)
    t.Log(rs1)
}
