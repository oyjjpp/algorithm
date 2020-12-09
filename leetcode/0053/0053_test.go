package leetcode

import (
    "testing"
)

func TestMaxSubArray(t *testing.T){
    data := []int{-2,1,-3,4,-1,2,1,-5,4}
    rs := maxSubArray(data)
    t.Log(rs)
}
